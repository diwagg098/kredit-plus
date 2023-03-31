package models

import "time"

type CreditTrasactionHistory struct {
	Id        int `gorm:"type:int;primary_key;autoIncrement:true" schema:"id"`
	UserId    int `gorm:"type:int" schema:"user_id"`
	LimitUsed int `gorm:"type:bigint" schema:"limit_used"`
	CreatedAt time.Time
	UpdateAt  time.Time
}

type CreditTrasactionHistoryResponse struct {
	CreditTrasaction CreditTrasaction
}
