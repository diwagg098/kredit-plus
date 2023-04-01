package controllers

import (
	"diwa/kredit-plus/pkg/models"
	"diwa/kredit-plus/pkg/utilities"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	UserId int
}

func (c *ctrl) Login(w http.ResponseWriter, r *http.Request) {
	var (
		ResponseLogin models.LoginResponse
	)

	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var dataPayload models.User
	err := decoder.Decode(&dataPayload)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error", err)
		Response(w, ctx, 500, false, INTERNALSERVER, err, nil)
		return
	}

	ctx, status, msg, data, err := c.uc.Login(dataPayload, ctx)

	if err != nil || status != 200 {
		Response(w, ctx, status, false, msg, nil, nil)
		return
	}

	timesExpired := time.Now().Add(time.Duration(24) * time.Hour).Unix() // expired date
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timesExpired,
		},
		UserId: data.Id,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	expirationTime := time.Now().Add(5 * time.Minute)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	if err != nil {
		Response(w, ctx, 500, true, INTERNALSERVER, nil, nil)
		return
	}

	ResponseLogin.AccessToken = tokenString
	ResponseLogin.TokenType = "bearer"
	ResponseLogin.User = data
	ResponseLogin.User.Password = ""

	Response(w, ctx, 200, true, SUCCESS, ResponseLogin, nil)
}

func (c *ctrl) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	r.ParseMultipartForm(32 << 20)

	file, fileHeader, err := r.FormFile("photo_ktp")
	fileName := ""
	if err == nil {
		result, err := utilities.UploadImage(ctx, file, fileHeader)
		if err != nil {
			ctx = utilities.Logf(ctx, "Error", err)
			Response(w, ctx, 500, false, INTERNALSERVER, err, nil)
			return
		}

		fileName = result
	}

	file, fileHeader, err = r.FormFile("photo_identity")
	fileNameKtp := ""
	if err == nil {
		result, err := utilities.UploadImage(ctx, file, fileHeader)
		if err != nil {
			ctx = utilities.Logf(ctx, "Error", err)
			Response(w, ctx, 500, false, INTERNALSERVER, err, nil)
			return
		}

		fileNameKtp = result
	}

	dataPayload := models.User{}
	err = decoder.Decode(&dataPayload, r.PostForm)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error", err)
		Response(w, ctx, 500, false, INTERNALSERVER, err, nil)
		return
	}

	dataPayload.PhotoIdentity = fileName
	dataPayload.PhotoKTP = fileNameKtp

	ctx, status, msg, data, err := c.uc.Register(dataPayload, ctx)
	if err != nil || status != 200 {
		Response(w, ctx, status, false, msg, nil, nil)
		return
	}
	Response(w, ctx, status, true, msg, data, nil)
}
