package models

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserCode  string     `json:"user_code"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	Email     string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"-"`
	RoleCode  string     `gorm:"type:varchar(5);not null" json:"role_code"`
	IsActive  bool       `gorm:"default:true" json:"is_active"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateUser struct {
	UserCode string `json:"user_code"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	RoleCode string `json:"role_code"`
	IsActive bool   `json:"is_active"`
}

type UserJoin struct {
	UserCode string `json:"user_code"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleCode string `json:"role_code"`
	RoleName string `json:"role_name"`
	IsActive bool   `son:"is_active"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	var lastUser Users
	tx.Order("id DESC").First(&lastUser)
	userCode := lastUser.UserCode
	re := regexp.MustCompile(`\d+`)
	numberStr := re.FindString(userCode)
	num, err := strconv.Atoi(numberStr)
	if err != nil {
		fmt.Println("ERR ", err)
		return
	}
	newID := num + 1
	u.UserCode = fmt.Sprintf("U%05d", newID)
	return nil
}
