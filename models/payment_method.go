package models

import (
	"time"

	"com.app/pos-app/utils"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PaymentMethodCode string    `gorm:"type:varchar(50);unique;not null" json:"payment_method_code"`
	Name              string    `gorm:"type:varchar(100);not null" json:"name"`
	Description       string    `gorm:"type:text" json:"description"`
	IsActive          bool      `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreatePaymentMethod struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (pm *PaymentMethod) BeforeCreate(tx *gorm.DB) (err error) {
	var last PaymentMethod
	tx.Order("id DESC").First(&last)
	paymentCode := last.PaymentMethodCode
	pm.PaymentMethodCode = utils.GenerateCode("PC", paymentCode, "5") //fmt.Sprintf("PRD%05d", newID) // PRD00001, PRD00002, ...
	return nil
}
