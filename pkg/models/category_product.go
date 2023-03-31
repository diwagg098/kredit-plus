package models

import "time"

type CategoryProduct struct {
	Id        int    `gorm:"type:int;primary_key;autoIncrement:true" schema:"id"`
	Name      string `gorm:"type:text" schema:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
