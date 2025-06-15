package controllers

import (
	"com.app/pos-app/utils"
	"github.com/gofiber/fiber/v2"
)

func GetRoles(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func GetOneRole(c *fiber.Ctx) error {
	return utils.Success(c, "Success getting data", fiber.Map{})
}
func CreateRole(c *fiber.Ctx) error {
	return utils.Success(c, "Success save data", fiber.Map{})
}
func UpdateRole(c *fiber.Ctx) error {
	return utils.Success(c, "Success update data", fiber.Map{})
}
func DeleteRole(c *fiber.Ctx) error {
	return utils.Success(c, "Success delete data", fiber.Map{})
}
