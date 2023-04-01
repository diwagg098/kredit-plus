package repository

import (
	"diwa/kredit-plus/pkg/models"
	"errors"
)

func (r *repo) CreateCreditTransaction(data models.CreditTrasaction) error {
	creditLimit := &models.LimitCredit{}

	tx := r.apps.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
	}

	// get detail limit user
	if err := tx.Model(&models.LimitCredit{}).First(&creditLimit, "user_id = ?", data.UserId).Error; err != nil {
		tx.Rollback()
	}

	creditLimit.Limit = creditLimit.Limit - data.TotalCredit
	if creditLimit.Limit < 0 {
		tx.Rollback()
		return errors.New("sufficient limit user")
	}

	// update limit user from limit used
	if err := tx.Model(&creditLimit).Updates(&creditLimit).Error; err != nil {
		tx.Rollback()
	}

	// add transaction history user
	if err := tx.Create(&models.CreditTrasactionHistory{
		UserId:    data.Id,
		LimitUsed: data.TotalCredit,
	}).Error; err != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}
