package controllers

import (
	"com.app/pos-app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetSaleItems(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func GetOneSaleItem(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func CreateSaleItem(c *fiber.Ctx) error {
	return utils.Success(c, "Success save data", fiber.Map{})
}
func UpdateSaleItem(c *fiber.Ctx) error {
	return utils.Success(c, "Success update data", fiber.Map{})
}
func DeleteSaleItem(c *fiber.Ctx) error {
	return utils.Success(c, "Success delete data", fiber.Map{})
}
