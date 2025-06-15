package models

import "time"

type Category struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryCode string     `gorm:"type:varchar(50);unique;not null" json:"category_code"`
	Name         string     `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt    *time.Time `gorm:"autoCreateTime" json:"created_at"`
}
