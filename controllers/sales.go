package controllers

import (
	"com.app/pos-app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetSales(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
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
