package models

import "time"

type LimitCredit struct {
	Id        int `gorm:"type:int;primary_key;autoIncrement:true" schema:"id"`
	UserId    int `gorm:"type:int" schema:"user_id"`
	Limit     int `gorm:"type:bigint" schema:"limit"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
