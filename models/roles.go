package models

import "time"

type Role struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleCode  string     `gorm:"type:varchar(50);unique;not null" json:"role_code"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
}
