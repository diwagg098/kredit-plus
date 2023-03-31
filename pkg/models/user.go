package models

import (
	"time"
)

type User struct {
	Id            int    `gorm:"type:int;primary_key;autoIncrement:true" schema:"id"`
	Nik           string `gorm:"type:text" schema:"nik" json:"nik"`
	FullName      string `gorm:"type:text" schema:"full_name" json:"full_name"`
	LegalName     string `gorm:"type:text" schema:"legal_name" json:"legal_name"`
	Pob           string `gorm:"type:text" schema:"pob" json:"pob"`
	Dob           string `gorm:"type:text" schema:"dob" json:"dob"`
	Salary        int    `gorm:"type:bigint" schema:"salary" json:"salary"`
	Username      string `gorm:"type:text" schema:"username" json:"username" validate:"required"`
	Password      string `gorm:"type:text" schema:"password" json:"password"`
	PhotoKTP      string `gorm:"type:text" schema:"photo_ktp" json:"photo_ktp"`
	PhotoIdentity string `gorm:"type:text" schema:"photo_identity" json:"photo_identity"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserResponse struct {
	ID            int    `json:"id"`
	Nik           string `json:"nik"`
	FullName      string `json:"full_name"`
	LegalName     string `json:"legal_name"`
	Pob           string `json:"pob"`
	Dob           string `json:"dob"`
	Salary        int    `json:"salary"`
	Username      string `json:"username"`
	PhotoKTP      string `json:"photo_ktp"`
	PhotoIdentity string `json:"photo_identity"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
