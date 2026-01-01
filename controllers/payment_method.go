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

func GetPaymentMethods(c *fiber.Ctx) error {
	keyword := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page", constants.Page))
	limit, _ := strconv.Atoi(c.Query("limit", constants.Limit))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	offset := (page - 1) * limit
	var paymentMethod []models.PaymentMethod
	query := database.DB.Select("*")
	if keyword != "" {
		query = query.Where("name ILIKE ?", keyword)
	}
	var total int64
	database.DB.Model(models.PaymentMethod{}).Count(&total)
	err := query.Limit(limit).Offset(offset).Find(&paymentMethod).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	return utils.Success(c, "Success getting data", fiber.Map{
		"result": paymentMethod,
		"pagination": fiber.Map{
			"current_page": page,
			"per_page":     limit,
			"total_data":   total,
			"total_pages":  int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}
func GetOnePaymentMethod(c *fiber.Ctx) error {
	code := c.Params("code")
	var paymentMethod models.PaymentMethod
	err := database.DB.Select("*").Where("id = ?", code).First(&paymentMethod).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	return utils.Success(c, "Success getting data", paymentMethod)
}
func CreatePaymentMethod(c *fiber.Ctx) error {
	validate := validator.New()
	var body models.CreatePaymentMethod
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
	var result models.PaymentMethod
	err := utils.WithTransaction(c.Context(), database.DB, func(tx *gorm.DB) error {
		payload := models.PaymentMethod{
			Name:        body.Name,
			Description: body.Description,
			IsActive:    body.IsActive,
		}
		err := tx.Create(&payload).Error
		if err != nil {
			return err
		}
		result = payload
		return nil
	})
	if err != nil {
		return utils.Error(c, 500, "Failed to save data", nil)
	}
	return utils.Success(c, "Success save data", result)
}
func UpdatePaymentMethod(c *fiber.Ctx) error {
	code := c.Params("code")
	var result models.PaymentMethod
	errFind := database.DB.Where("payment_method_code = ?", code).First(&result).Error
	if errFind != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	var body models.CreatePaymentMethod
	errBody := c.BodyParser(&body)
	if errBody != nil {
		return utils.Error(c, 400, "Invalid request", nil)
	}
	err := utils.WithTransaction(c.Context(), database.DB, func(tx *gorm.DB) error {
		payload := models.PaymentMethod{
			Name:        body.Name,
			Description: body.Description,
			IsActive:    body.IsActive,
		}
		errUpdate := tx.Where("payment_method_code = ?", code).Updates(&payload).Error
		if errUpdate != nil {
			return errUpdate
		}
		result.Name = body.Name
		result.Description = body.Description
		result.IsActive = body.IsActive
		return nil
	})
	if err != nil {
		return utils.Error(c, 500, "Failed update data", nil)
	}
	return utils.Success(c, "Success update data", result)
}
func DeletePaymentMethod(c *fiber.Ctx) error {
	code := c.Params("code")
	var result models.PaymentMethod
	err := utils.WithTransaction(c.Context(), database.DB, func(tx *gorm.DB) error {
		errFind := tx.Where("payment_method_code = ?", code).First(&result).Error
		if errFind != nil {
			return utils.Error(c, 404, "Data not found", nil)
		}
		err := tx.Where("payment_method_code = ?", code).Delete(&models.PaymentMethod{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return utils.Error(c, 404, "Failed to delete data", nil)
	}
	return utils.Success(c, "Success delete data", result)
}
