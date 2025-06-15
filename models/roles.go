package models

import (
	"time"

	"com.app/pos-app/utils"
	"gorm.io/gorm"
)

type Role struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleCode  string     `gorm:"type:varchar(50);unique;not null" json:"role_code"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateRole struct {
	RoleCode string `json:"role_code"`
	Name     string `json:"name" validate:"required,min=3"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	var last Role
	tx.Order("id DESC").First(&last)
	roleCode := last.RoleCode
	r.RoleCode = utils.GenerateCode("RL", roleCode, "3")
	return nil
}
