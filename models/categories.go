package models

import (
	"time"

	"com.app/pos-app/utils"
	"gorm.io/gorm"
)

type Category struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryCode string     `gorm:"type:varchar(50);unique;not null" json:"category_code"`
	Name         string     `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt    *time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateCategory struct {
	Name string `json:"name" validate:"required"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	var last Category
	tx.Order("id DESC").First(&last)
	categoryCode := last.CategoryCode
	c.CategoryCode = utils.GenerateCode("C", categoryCode, "5") //fmt.Sprintf("PRD%05d", newID) // PRD00001, PRD00002, ...
	return nil
}
