package controllers

import (
	"math"
	"strconv"

	"com.app/pos-app/constants"
	"com.app/pos-app/database"
	"com.app/pos-app/models"
	"com.app/pos-app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	code := c.Params("code")
	var result models.Sales
	err := database.DB.Where("sales_code = ?", code).First(&result).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	return utils.Success(c, "Success getting data", result)
}
func CreateSale(c *fiber.Ctx) error {
	validate := validator.New()
	var body models.CreateSales
	errBody := c.BodyParser(&body)
	if errBody != nil {
		return utils.Error(c, 400, "Invalid request", nil)
	}
	errValidate := validate.Struct(body)
	if errValidate != nil {
		for _, e := range errValidate.(validator.ValidationErrors) {
			errorField := utils.ValidatorForm(e)
			return utils.Error(c, 400, "Validation error", errorField)
		}
	}
	var result models.Sales
	err := utils.WithTransaction(c.Context(), database.DB, func(tx *gorm.DB) error {
		payload := models.Sales{
			UserCode:          body.UserCode,
			PaymentMethodCode: body.PaymentMethodCode,
			Total:             body.Total,
			Discount:          body.Discount,
			FinalTotal:        body.FinalTotal,
		}
		errInsert := tx.Create(&payload).Error
		if errInsert != nil {
			return errInsert
		}
		result.UserCode = body.UserCode
		result.Total = body.Total
		result.Discount = body.Discount
		result.FinalTotal = body.FinalTotal
		return nil
	})
	if err != nil {
		return utils.Error(c, 500, "Failed to save data", nil)
	}
	return utils.Success(c, "Success save data", result)
}
func UpdateSale(c *fiber.Ctx) error {
	return utils.Success(c, "Success update data", fiber.Map{})
}
func DeleteSale(c *fiber.Ctx) error {
	code := c.Params("code")
	var result models.Sales
	err := utils.WithTransaction(c.Context(), database.DB, func(tx *gorm.DB) error {
		errFind := tx.Where("sales_code = ?", code).First(&result).Error
		if errFind != nil {
			return utils.Error(c, 404, "Data not found", nil)
		}
		errDelete := tx.Where("sales_code = ?", code).Delete(&result).Error
		if errDelete != nil {
			return errDelete
		}
		return nil
	})
	if err != nil {
		return utils.Error(c, 404, "Failed to delete data", nil)
	}
	return utils.Success(c, "Success delete data", result)
}
