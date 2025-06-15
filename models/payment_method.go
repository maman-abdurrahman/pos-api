package models

import "time"

type PaymentMethod struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PaymentMethodCode string    `gorm:"type:varchar(50);unique;not null" json:"payment_method_code"`
	Name              string    `gorm:"type:varchar(100);not null" json:"name"`
	Description       string    `gorm:"type:text" json:"description"`
	IsActive          bool      `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
}
