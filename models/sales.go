package models

import (
	"time"

	"com.app/pos-app/utils"
	"gorm.io/gorm"
)

type Sales struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SalesCode         string    `gorm:"type:varchar(50);unique;not null" json:"sales_code"`
	InvoiceNumber     string    `gorm:"type:varchar(100);unique;not null" json:"invoice_number"`
	UserCode          string    `gorm:"not null" json:"user_code"`
	PaymentMethodCode string    `gorm:"type:varchar(50);not null" json:"payment_method_code"`
	Total             float64   `gorm:"type:numeric(12,2);not null" json:"total"`
	Discount          float64   `gorm:"type:numeric(12,2);default:0" json:"discount"`
	FinalTotal        float64   `gorm:"type:numeric(12,2);not null" json:"final_total"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateSales struct {
	UserCode          string  `json:"user_code"`
	PaymentMethodCode string  `json:"payment_method_code" validate:"required"`
	Total             float64 `json:"total"`
	Discount          float64 `json:"discount"`
	FinalTotal        float64 `json:"final_total"`
}

func (s *Sales) BeforeCreate(tx *gorm.DB) (err error) {
	var lastSales Sales
	tx.Order("id DESC").First(&lastSales)
	salesCode := lastSales.SalesCode
	s.SalesCode = utils.GenerateCode("S", salesCode, "5")
	return nil
}
