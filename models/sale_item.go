package models

import (
	"com.app/pos-app/utils"
	"gorm.io/gorm"
)

type SaleItem struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	SaleCode    string  `gorm:"type:varchar(50);not null" json:"sale_code"`    // FK ke Sales.SalesCode
	ProductCode string  `gorm:"type:varchar(50);not null" json:"product_code"` // FK ke Product.ProductCode
	Quantity    int     `gorm:"not null" json:"quantity"`
	UnitPrice   float64 `gorm:"type:numeric(12,2);not null" json:"unit_price"`
	Subtotal    float64 `gorm:"type:numeric(12,2);not null" json:"subtotal"`
}

type CreateSaleItem struct {
	SaleCode    string  `json:"sale_code"`
	ProductCode string  `json:"product_code"`
	Quantity    int     `json:"quantity" validate:"required, gt=1"`
	UnitPrice   float64 `json:"unit_price" validate:"required, gt=0"`
	Subtotal    float64 `json:"subtotal" validate:"required, gt=0"`
}

func (si *SaleItem) BeforeCreate(tx *gorm.DB) (err error) {
	var last SaleItem
	tx.Order("id DESC").First(&last)
	saleCode := last.SaleCode
	si.SaleCode = utils.GenerateCode("SI", saleCode, "5") //fmt.Sprintf("PRD%05d", newID) // PRD00001, PRD00002, ...
	return nil
}
