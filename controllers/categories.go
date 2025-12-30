package controllers

import (
	"math"
	"strconv"

	"com.app/pos-app/database"
	"com.app/pos-app/models"
	"com.app/pos-app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	keyword := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "2"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	query := database.DB.
		Table("categories").
		Select("*")
	if keyword != "" {
		query = query.Where(
			"categories.name ILIKE ? OR categories.category_code ILIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%",
		)
	}
	var total int64
	offset := (page - 1) * limit
	var category []models.Category
	database.DB.Model(models.Category{}).Count(&total)
	err := query.
		Limit(limit).
		Offset(offset).
		Find(&category).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	return utils.Success(c, "Success getting data 1", fiber.Map{
		"result": category,
		"pagination": fiber.Map{
			"current_page": page,
			"per_page":     limit,
			"total_data":   total,
			"total_pages":  int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}
func GetOneCategory(c *fiber.Ctx) error {
	code := c.Params("code")
	var category []models.Category
	err := database.DB.Table("categories").Select("*").Where("category_code = ?", code).First(&category).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	return utils.Success(c, "Success getting data", category)
}
func CreateCategory(c *fiber.Ctx) error {
	validate := validator.New()
	var body models.CreateCategory
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
	payload := models.Category{
		Name: body.Name,
	}
	errStore := database.DB.Create(&payload).Error
	if errStore != nil {
		return utils.Error(c, 500, "Failed to save data", nil)
	}
	return utils.Success(c, "Success save data", payload)
}
func UpdateCategory(c *fiber.Ctx) error {
	code := c.Params("code")
	var category models.Category
	errFindCategory := database.DB.Where("category_code = ?", code).First(&category).Error
	if errFindCategory != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	var body models.CreateCategory
	errBody := c.BodyParser(&body)
	if errBody != nil {
		return utils.Error(c, 400, "Invalid request", nil)
	}
	payload := models.Category{
		Name: body.Name,
	}
	errUpdate := database.DB.Model(models.Category{}).Where("category_code = ?", code).Updates(payload).Error
	if errUpdate != nil {
		return utils.Error(c, 500, "Failed update data", nil)
	}
	result := category
	result.Name = body.Name
	return utils.Success(c, "Success update data", result)
}
func DeleteCategory(c *fiber.Ctx) error {
	code := c.Params("code")
	var category models.Category
	errFind := database.DB.Where("category_code = ?", code).First(&category).Error
	if errFind != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	errDelete := database.DB.Where("category_code = ?", code).Delete(&models.Category{}).Error
	if errDelete != nil {
		return utils.Error(c, 404, "Failed to delete data", nil)
	}
	return utils.Success(c, "Success delete data", category)
}
