package routes

import (
	"com.app/pos-app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", controllers.GetUsers)
	api.Get("/users/:code", controllers.GetOneUser)
	api.Post("/users/create", controllers.PostUser)
	api.Put("/users/update/:code", controllers.UpdateUser)
	api.Delete("/users/delete/:code", controllers.DeleteUser)
	api.Get("/products", controllers.GetProducts)
	api.Get("/products/:code", controllers.GetOneProducts)
	api.Post("/products/create", controllers.CreateProducts)
	api.Put("/products/update/:code", controllers.UpdateProduct)
	api.Delete("/products/delete/:code", controllers.DeleteProduct)
}
