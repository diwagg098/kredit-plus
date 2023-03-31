package repository

import (
	"database/sql"
	"diwa/kredit-plus/pkg/models"

	"gorm.io/gorm"
)

type repo struct {
	apps *gorm.DB
}

type Repo interface {
	FindAllWithWhere(i interface{}, where map[string]interface{}) error
	FindOne(i interface{}, where map[string]interface{}) (error, int64)
	InsertData(i interface{}) error
	FindOneWithTableName(i interface{}, where map[string]interface{}, tablename string) error
	DinamicFindQueryRaw(i interface{}, query string) (*sql.Rows, error)
	DeleteData(i interface{}, where map[string]interface{}) error
	UpdateData(i interface{}, where map[string]interface{}, data map[string]interface{}) error
	FindAll(i interface{}) error

	FindAllWithWhereV2(i interface{}, cond ...interface{}) error

	FindAllWithWhereV3(i interface{}, where string) error

	// implement concurrent transaction
	CreateCreditTransaction(data models.CreditTrasaction) error
}

func NewRepo(apps *gorm.DB) Repo {
	return &repo{apps}
}
