package controllers

import (
	"math"
	"strconv"

	"com.app/pos-app/database"
	"com.app/pos-app/models"
	"com.app/pos-app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetRoles(c *fiber.Ctx) error {
	keyword := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "2"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 2
	}
	offset := (page - 1) * limit

	query := database.DB.Table("roles").Select("roles.*")
	if keyword != "" {
		query = query.Where("name ILIKE ?", "%"+keyword+"%")
	}
	var roles []models.Role
	err := query.
		Limit(limit).
		Offset(offset).
		Find(&roles).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	var total int64
	database.DB.Model(models.Role{}).Count(&total)

	return utils.Success(c, "Success getting data", fiber.Map{
		"result": roles,
		"pagination": fiber.Map{
			"current_page": page,
			"per_page":     limit,
			"total_data":   total,
			"total_pages":  int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}
func GetOneRole(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func CreateRole(c *fiber.Ctx) error {
	var roleBody models.CreateRole
	errBody := c.BodyParser(&roleBody)
	if errBody != nil {
		return utils.Error(c, 500, "Invalid request", nil)
	}
	role := models.Role{
		Name: roleBody.Name,
	}
	err := database.DB.Create(&role).Error
	if err != nil {
		return utils.Error(c, 500, "Failed to save data", nil)
	}
	return utils.Success(c, "Success save data", role)
}
func UpdateRole(c *fiber.Ctx) error {
	return utils.Success(c, "Success update data", fiber.Map{})
}
func DeleteRole(c *fiber.Ctx) error {
	return utils.Success(c, "Success delete data", fiber.Map{})
}
