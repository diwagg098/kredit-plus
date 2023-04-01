package usecase

import (
	"context"
	"diwa/kredit-plus/pkg/models"
	"diwa/kredit-plus/pkg/utilities"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func (uc *uc) Login(data models.User, ctx context.Context) (context.Context, int, string, models.User, error) {
	where := map[string]interface{}{
		"username": data.Username,
	}

	realPW := data.Password

	err, _ := uc.query.FindOne(&data, where)
	if err != nil {
		ctx = utilities.Logf(ctx, "Get Data User Not Found -> : %v", err)
		return ctx, 404, "User not Found", data, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(realPW))
	if err != nil {
		ctx = utilities.Logf(ctx, "Data user not found -> : %v", err)
		return ctx, 404, "User not found", data, err
	}

	return ctx, 200, "Success", data, err
}

func (uc *uc) Register(data models.User, ctx context.Context) (context.Context, int, string, models.User, error) {
	var (
		dataUser      []models.User
		percent_limit = os.Getenv("limit")
	)

	where := map[string]interface{}{
		"username": data.Username,
	}

	err := uc.query.FindAllWithWhere(&dataUser, where)
	if err != nil {
		ctx = utilities.Logf(ctx, "Validation Error execute : %v", err)
		return ctx, 500, "User not Found", data, err
	}

	if len(dataUser) > 0 {
		ctx = utilities.Logf(ctx, "Username is using", data, err)
		return ctx, 409, "username is using", data, err
	}

	password := []byte(data.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		ctx = utilities.Logf(ctx, "Internal server Error -> : %v", err)
		return ctx, 500, "Internal Server Error ", data, err
	}

	data.Password = string(hashedPassword)

	err = uc.query.InsertData(&data)
	if err != nil {
		ctx = utilities.Logf(ctx, "Internal server Error -> : %v", err)
		return ctx, 500, "Internal Server Error ", data, err
	}

	// generate limit and tenor
	if data.Id != 0 {
		// first limit credit = salary * 70%
		limit_percent_int, _ := strconv.Atoi(percent_limit)
		limit := data.Salary * limit_percent_int / 100

		dataLimit := models.LimitCredit{
			UserId: data.Id,
			Limit:  limit,
		}

		err = uc.query.InsertData(&dataLimit)
		if err != nil {
			ctx = utilities.Logf(ctx, "Internal server Error -> : %v", err)
			return ctx, 500, "Internal Server Error", data, err
		}

		tenor1 := limit * 25 / 100
		tenor2 := limit * 50 / 100
		tenor3 := limit * 75 / 100
		tenor4 := limit * 100 / 100
		dataCredit := models.LimitTenor{
			Tenor1: tenor1,
			Tenor2: tenor2,
			Tenor3: tenor3,
			Tenor4: tenor4,
			UserId: data.Id,
		}

		err = uc.query.InsertData(&dataCredit)
		if err != nil {
			ctx = utilities.Logf(ctx, "Internal server Error -> : %v", err)
			return ctx, 500, "Internal Server Error", data, err
		}
	}
	return ctx, 200, "Success Add Account", data, nil
}

func (uc *uc) GetDataListUser(ctx context.Context) ([]models.UserResponseList, error, context.Context) {
	// dataCTX := utilities.GetDataCTX(ctx)
	var data []models.UserResponseList

	query := "SELECT * FROM users"
	_, err := uc.query.DinamicFindQueryRaw(&data, query)
	if err != nil {
		ctx = utilities.Logf(ctx, "error get data customer -> : %v", err)
		return data, err, ctx
	}

	return data, nil, ctx
}

func (uc *uc) GetUserById(id int, ctx context.Context) (context.Context, *models.UserResponse, error) {
	var data models.User
	var res models.UserResponse
	var limit models.LimitCredit
	var tenor models.LimitTenor

	where := map[string]interface{}{
		"id": id,
	}

	whereUserId := map[string]interface{}{
		"user_id": id,
	}

	err, _ := uc.query.FindOne(&data, where)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error query find user by id : %v", err)
		return ctx, nil, err
	}

	err, _ = uc.query.FindOne(&limit, whereUserId)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error query find user by id : %v", err)
		return ctx, nil, err
	}

	err, _ = uc.query.FindOne(&tenor, whereUserId)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error query find user by id : %v", err)
		return ctx, nil, err
	}

	data.Password = ""

	res.User = data
	res.Limit = limit.Limit
	res.Tenor = tenor

	log.Println(res)

	return ctx, &res, nil
}
