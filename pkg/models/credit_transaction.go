package models

import "time"

type CreditTrasaction struct {
	Id                int    `gorm:"type:int;primary_key;autoIncrement:true" schema:"id" json:"id"`
	ContractId        string `gorm:"type:text;unique" schema:"contract_id" json:"contract_id"`
	AssetName         string `gorm:"type:text" schema:"asset_name" json:"asset_name"`
	UserId            int    `gorm:"type:int" schema:"user_id" json:"user_id"`
	Tenor             int    `gorm:"type:bigint" schema:"tenor" json:"tenor"`
	ProductCategoryId int    `gorm:"type:int" schema:"product_category_id" json:"product_category_id"`
	OTR               int    `gorm:"type:bigint" schema:"otr" json:"otr"`
	AdminFee          int    `gorm:"type:bigint" schema:"admin_fee" json:"admin_fee"`
	TotalCredit       int    `gorm:"type:bigint" schema:"total_credit" json:"total_credit"`
	TotalBunga        int    `gorm:"type:bigint" schema:"total_bunga" json:"total_bunga"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
