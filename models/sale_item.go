package models

type SaleItem struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	SaleCode    string  `gorm:"type:varchar(50);not null" json:"sale_code"`    // FK ke Sales.SalesCode
	ProductCode string  `gorm:"type:varchar(50);not null" json:"product_code"` // FK ke Product.ProductCode
	Quantity    int     `gorm:"not null" json:"quantity"`
	UnitPrice   float64 `gorm:"type:numeric(12,2);not null" json:"unit_price"`
	Subtotal    float64 `gorm:"type:numeric(12,2);not null" json:"subtotal"`
}
