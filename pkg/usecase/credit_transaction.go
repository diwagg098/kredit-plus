package usecase

import (
	"context"
	"diwa/kredit-plus/pkg/models"
	"diwa/kredit-plus/pkg/utilities"
)

func (uc *uc) CreateTransaction(data models.CreditTrasaction, ctx context.Context) (context.Context, int, string, models.CreditTrasaction, error) {
	err := uc.query.CreateCreditTransaction(data)
	if err != nil {
		ctx = utilities.Logf(ctx, "error insert data credit transaction : %v", err)
		return ctx, 500, "internal server error", data, err
	}

	return ctx, 200, "Transaction Success", data, nil
}

func (uc *uc) CreditTransactionList(ctx context.Context) ([]models.CreditTrasaction, error, context.Context) {
	var (
		data []models.CreditTrasaction
	)

	query := "select * from credit_trasactions"
	_, err := uc.query.DinamicFindQueryRaw(&data, query)
	if err != nil {
		ctx = utilities.Logf(ctx, "error get data credit transaction : %v", err)
		return data, err, ctx
	}

	return data, nil, ctx
}

func (uc *uc) FindByIdCredit(id int, ctx context.Context) (context.Context, models.CreditTrasaction, error) {
	var data models.CreditTrasaction

	where := map[string]interface{}{
		"id": id,
	}

	err, count := uc.query.FindOne(&data, where)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error query find credit transaction by id : %v", err)
		return ctx, data, err
	}

	if count == 0 {
		ctx = utilities.Logf(ctx, "no sql row : %v", err)
		return ctx, data, err
	}

	return ctx, data, nil
}

func (uc *uc) FindByIdUser(id int, ctx context.Context) (context.Context, models.CreditTrasaction, error) {
	var data models.CreditTrasaction

	where := map[string]interface{}{
		"user_id": id,
	}

	err, count := uc.query.FindOne(&data, where)
	if err != nil || count == 0 {
		ctx = utilities.Logf(ctx, "Error query find credit transaction by id : %v", err)
		return ctx, data, err
	}

	return ctx, data, nil
}
