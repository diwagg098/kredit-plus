package models

import "time"

type LimitTenor struct {
	Id        int `gorm:"type:int;primary_key;autoIncrement:true" schema:"id"`
	UserId    int `gorm:"type:int" schema:"user_id"`
	Tenor1    int `gorm:"type:bigint" schema:"tenor_1"`
	Tenor2    int `gorm:"type:bigint" schema:"tenor_2"`
	Tenor3    int `gorm:"type:bigint" schema:"tenor_3"`
	Tenor4    int `gorm:"type:bigint" schema:"tenor_4"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
