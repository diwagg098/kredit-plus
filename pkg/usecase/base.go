package usecase

import (
	"context"
	"diwa/kredit-plus/pkg/models"
	"diwa/kredit-plus/pkg/repository"
)

type uc struct {
	query repository.Repo
}

type UC interface {
	Login(data models.User, ctx context.Context) (context.Context, int, string, models.User, error)
	Register(data models.User, ctx context.Context) (context.Context, int, string, models.User, error)
	GetDataListUser(ctx context.Context) ([]models.UserResponse, error, context.Context)
	GetUserById(id int, ctx context.Context) (context.Context, *models.User, error)
	CreateTransaction(data models.CreditTrasaction, ctx context.Context) (context.Context, int, string, models.CreditTrasaction, error)
}

func NewUC(r repository.Repo) UC {
	return &uc{query: r}
}
