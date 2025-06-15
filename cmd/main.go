package main

import (
	"com.app/pos-app/config"
	"com.app/pos-app/database"
	"com.app/pos-app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config := config.GetAppConfig()
	database.ConnectionDB()
	routes.PublicRoutes(app)
	app.Listen(config.Server.Hostname + ":" + config.Server.Port)
}
