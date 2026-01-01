package controllers

import (
	"math"
	"strconv"

	"com.app/pos-app/constants"
	"com.app/pos-app/database"
	"com.app/pos-app/models"
	"com.app/pos-app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetSales(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.Query("page", constants.Page))
	limit, _ := strconv.Atoi(c.Query("limit", constants.Limit))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	offset := (page - 1) * limit
	var result []models.Sales
	query := database.DB.Select("*")
	if keyword != "" {
		query = query.Where(
			"sales_code ILIKE ? OR invoice_number ILIKE ? ",
			"%"+keyword+"%",
			"%"+keyword+"%",
		)
	}
	var total int64
	database.DB.Model(models.Sales{}).Count(&total)
	err := query.Limit(limit).Offset(offset).Find(&result).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	return utils.Success(c, "Success getting data", fiber.Map{
		"result": result,
		"pagination": fiber.Map{
			"current_page": page,
			"per_page":     limit,
			"total_data":   total,
			"total_pages":  int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}
func GetOneSale(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func CreateSale(c *fiber.Ctx) error {
	return utils.Success(c, "Success save data", fiber.Map{})
}
func UpdateSale(c *fiber.Ctx) error {
	return utils.Success(c, "Success update data", fiber.Map{})
}
func DeleteSale(c *fiber.Ctx) error {
	return utils.Success(c, "Success delete data", fiber.Map{})
}
