package models

import (
	"time"

	"com.app/pos-app/utils"
	"gorm.io/gorm"
)

type Product struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	ProductCode  string     `gorm:"unique;not null" json:"product_code"`
	Name         string     `gorm:"size:100;not null" json:"name"`
	CategoryCode string     `json:"category_code"`
	Price        float64    `gorm:"not null" json:"price"`
	Stock        int        `gorm:"not null" json:"stock"`
	DateIn       *time.Time `json:"date_in"`
	CreatedBy    string     `json:"created_by"`
	UpdatedBy    string     `json:"updated_by"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

// type CreateProduct struct {
// 	ProductCode  string  `json:"product_code"`
// 	Name         string  `json:"name"`
// 	CategoryCode string  `json:"category_code"`
// 	Price        float64 `json:"price"`
// 	Stock        int     `json:"stock"`
// }

type CreateProduct struct {
	ProductCode  string  `json:"product_code"`
	Name         string  `json:"name" validate:"required,min=3"`
	CategoryCode string  `json:"category_code" validate:"required"`
	Price        float64 `json:"price" validate:"required,gt=0"`
	Stock        int     `json:"stock" validate:"required,gte=0"`
}

type ProductWithCategory struct {
	ProductCode  string  `json:"product_code"`
	ProductName  string  `json:"product_name"`
	CategoryCode string  `json:"category_code"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryName string  `json:"category_name"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	var last Product
	tx.Order("id DESC").First(&last)
	productCode := last.ProductCode
	p.ProductCode = utils.GenerateCode("PRD", productCode, "5") //fmt.Sprintf("PRD%05d", newID) // PRD00001, PRD00002, ...
	return nil
}
