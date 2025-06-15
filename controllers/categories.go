package controllers

import (
	"com.app/pos-app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func GetOneCategory(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func CreateCategory(c *fiber.Ctx) error {
	return utils.Success(c, "Success save data", fiber.Map{})
}
func UpdateCategory(c *fiber.Ctx) error {
	return utils.Success(c, "Success update data", fiber.Map{})
}
func DeleteCategory(c *fiber.Ctx) error {
	return utils.Success(c, "Success delete data", fiber.Map{})
}
