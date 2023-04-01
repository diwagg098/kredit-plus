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

type UserResponseList struct {
	Id            int    `gorm:"type:int;primary_key;autoIncrement:true" schema:"id"`
	Nik           string `gorm:"type:text" schema:"nik" json:"nik"`
	FullName      string `gorm:"type:text" schema:"full_name" json:"full_name"`
	LegalName     string `gorm:"type:text" schema:"legal_name" json:"legal_name"`
	Pob           string `gorm:"type:text" schema:"pob" json:"pob"`
	Dob           string `gorm:"type:text" schema:"dob" json:"dob"`
	Salary        int    `gorm:"type:bigint" schema:"salary" json:"salary"`
	Username      string `gorm:"type:text" schema:"username" json:"username" validate:"required"`
	PhotoKTP      string `gorm:"type:text" schema:"photo_ktp" json:"photo_ktp"`
	PhotoIdentity string `gorm:"type:text" schema:"photo_identity" json:"photo_identity"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserResponse struct {
	User  User       `json:"user_detail"`
	Limit int        `json:"limit"`
	Tenor LimitTenor `json:"limit_tenor"`
}
