package controllers

import (
	"com.app/pos-app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetPaymentMethods(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func GetOnePaymentMethod(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func CreatePaymentMethod(c *fiber.Ctx) error {
	return utils.Success(c, "Success save data", fiber.Map{})
}
func UpdatePaymentMethod(c *fiber.Ctx) error {
	return utils.Success(c, "Success update data", fiber.Map{})
}
func DeletePaymentMethod(c *fiber.Ctx) error {
	return utils.Success(c, "Success delete data", fiber.Map{})
}
