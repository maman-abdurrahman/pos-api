package controllers

import (
	"math"
	"strconv"

	"com.app/pos-app/database"
	"com.app/pos-app/models"
	"com.app/pos-app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	keyword := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "2"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	offset := (page - 1) * limit
	var users []models.UserJoin
	query := database.DB.
		Table("users").
		Select("users.user_code, users.name, users.email, users.is_active, users.role_code, roles.name AS role_name").
		Joins("LEFT JOIN roles ON roles.role_code = users.role_code")
	if keyword != "" {
		query = query.Where("users.name ILIKE ?", "%"+keyword+"%")
	}
	var total int64
	database.DB.Model(models.Users{}).Count(&total)
	err := query.
		Limit(limit).
		Offset(offset).
		Find(&users).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found")
	}
	return utils.Success(c, "Success getting data", fiber.Map{
		"result": users,
		"pagination": fiber.Map{
			"current_page": page,
			"per_page":     limit,
			"total_data":   total,
			"total_pages":  int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}

func GetOneUser(c *fiber.Ctx) error {
	code := c.Params("code")
	var user models.UserJoin
	err := database.DB.Table("users").
		Select("users.user_code, users.name, users.email, users.is_active, users.role_code, roles.name AS role_name").
		Joins("LEFT JOIN roles ON roles.role_code = users.role_code").
		Where("user_code = ?", code).
		First(&user).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found")
	}
	return utils.Success(c, "Success getting data", user)
}

func PostUser(c *fiber.Ctx) error {
	var userBody models.CreateUser
	errBody := c.BodyParser(&userBody)
	if errBody != nil {
		return utils.Error(c, 400, "Invalid request")
	}
	if len(userBody.Password) < 8 {
		return utils.Error(c, 500, "Password to small min 8 character")
	}
	hashPass, errHash := utils.HashPassword(userBody.Password)
	if errHash != nil {
		return utils.Error(c, 500, "Error hash password")
	}
	userPayload := models.Users{
		Name:     userBody.Name,
		Email:    userBody.Email,
		Password: hashPass,
		RoleCode: userBody.RoleCode,
		IsActive: userBody.IsActive,
	}
	errStore := database.DB.Create(&userPayload).Error
	if errStore != nil {
		return utils.Error(c, 500, "Failed to save data")
	}
	return utils.Success(c, "Success save data", userPayload)
}

func UpdateUser(c *fiber.Ctx) error {
	code := c.Params("code")
	var users models.Users
	errFindUser := database.DB.Where("user_code = ?", code).First(&users).Error
	if errFindUser != nil {
		return utils.Error(c, 404, "Data not found")
	}
	var userBody models.CreateUser
	errBody := c.BodyParser(&userBody)
	if errBody != nil {
		return utils.Error(c, 400, "Invalid request")
	}
	userPayload := models.Users{
		Name:     userBody.Name,
		Email:    userBody.Email,
		Password: userBody.Password,
		RoleCode: userBody.RoleCode,
		IsActive: userBody.IsActive,
	}
	errUpdate := database.DB.Model(models.Users{}).
		Where("user_code =?", code).
		Updates(userPayload).Error
	if errUpdate != nil {
		return utils.Error(c, 500, "Failed update data")
	}
	return utils.Success(c, "Success update data", userPayload)
}

func DeleteUser(c *fiber.Ctx) error {
	code := c.Params("code")
	var users models.Users
	errFind := database.DB.Where("user_code = ?", code).First(&users).Error
	if errFind != nil {
		return utils.Error(c, 404, "Data not found")
	}
	errDelete := database.DB.Where("user_code = ?", code).Delete(&models.Users{}).Error
	if errDelete != nil {
		return utils.Error(c, 404, "Failed to delete data")
	}
	return utils.Success(c, "Success delete data", users)
}
