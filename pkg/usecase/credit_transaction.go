package usecase

import (
	"context"
	"diwa/kredit-plus/pkg/models"
	"diwa/kredit-plus/pkg/utilities"
	"log"
)

func (uc *uc) CreateTransaction(data models.CreditTrasaction, ctx context.Context) (context.Context, int, string, models.CreditTrasaction, error) {
	log.Println(data)
	err := uc.query.CreateCreditTransaction(data)
	if err != nil {
		ctx = utilities.Logf(ctx, "error insert data action : %v", err)
		return ctx, 500, "internal server error", data, err
	}

	return ctx, 200, "Transaction Success", data, nil
}
