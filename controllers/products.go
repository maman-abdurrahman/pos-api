package controllers

import (
	"database/sql"
	"log"
	"math"
	"strconv"
	"time"

	"com.app/pos-app/database"
	"com.app/pos-app/models"
	"com.app/pos-app/utils"
	"github.com/gofiber/fiber/v2"
)

var DB *sql.DB

func GetProducts(c *fiber.Ctx) error {
	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "2"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	var products []models.ProductWithCategory
	query := database.DB.
		Table("products").
		Select("products.product_code, products.name AS product_name, products.category_code, products.price, products.stock, categories.name AS category_name").
		Joins("LEFT JOIN categories ON products.category_code = categories.category_code")
	if search != "" {
		query = query.Where("products.name ILIKE ? OR categories.name ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	err := query.
		Limit(limit).
		Offset(offset).
		Scan(&products).Error
	if err != nil {
		return utils.Error(c, 500, "Failed to fetch users", nil)
	}
	var total int64
	database.DB.Model(models.Product{}).Count(&total)
	return utils.Success(c, "Success get data", fiber.Map{
		"result": products,
		"pagination": fiber.Map{
			"current_page": page,
			"per_page":     limit,
			"total_data":   total,
			"total_pages":  int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}

func GetOneProducts(c *fiber.Ctx) error {
	code := c.Params("code")

	var products models.ProductWithCategory
	err := database.DB.
		Table("products").
		Select("products.product_code, products.name AS product_name, products.category_code, products.price, products.stock, categories.name AS category_name").
		Joins("LEFT JOIN categories ON products.category_code = categories.category_code").
		Where("product_code = ?", code).
		First(&products).Error
	if err != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	return utils.Success(c, "Success get data", products)
}

func CreateProducts(c *fiber.Ctx) error {
	var request models.CreateProduct

	errReq := c.BodyParser(&request)
	if errReq != nil {
		return utils.Error(c, 400, "Invalid request", nil)
	}

	product := models.Product{
		ProductCode:  request.ProductCode,
		Name:         request.Name,
		CategoryCode: request.CategoryCode,
		Price:        request.Price,
		Stock:        request.Stock,
	}

	err := database.DB.Create(&product).Error
	if err != nil {
		return utils.Error(c, 500, "Failed to insert product", nil)
	}
	return utils.Success(c, "Success save data", product)
}

func UpdateProduct(c *fiber.Ctx) error {
	code := c.Params("code")
	var product models.Product
	errFindProduct := database.DB.Where("product_code = ?", code).First(&product).Error
	if errFindProduct != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	var bodyProduct models.CreateProduct
	errBody := c.BodyParser(&bodyProduct)
	if errBody != nil {
		return utils.Error(c, 400, "Invalid request", nil)
	}
	now := time.Now()
	productStore := models.Product{
		ProductCode:  bodyProduct.ProductCode,
		Name:         bodyProduct.Name,
		CategoryCode: bodyProduct.CategoryCode,
		Price:        bodyProduct.Price,
		Stock:        bodyProduct.Stock,
		DateIn:       &now,
	}
	errStore := database.DB.Model(&models.Product{}).
		Where("product_code = ?", code).
		Updates(productStore).Error
	if errStore != nil {
		return utils.Error(c, 500, "Failed to update product", nil)
	}
	return utils.Success(c, "Success update data", productStore)
}

func DeleteProduct(c *fiber.Ctx) error {
	code := c.Params("code")
	log.Println("PARAMS ", code)
	var product models.Product
	errFindProduct := database.DB.Where("product_code = ?", code).First(&product).Error
	if errFindProduct != nil {
		return utils.Error(c, 404, "Data not found", nil)
	}
	errDelete := database.DB.Where("product_code = ?", code).Delete(&models.Product{}).Error
	if errDelete != nil {
		return utils.Error(c, 500, "Failed to delete product", nil)
	}
	return utils.Success(c, "Success delete data", product)
}
