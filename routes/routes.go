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

	api.Get("/roles", controllers.GetRoles)
	api.Get("/roles/:code", controllers.GetOneRole)
	api.Post("/roles/create", controllers.CreateRole)
	api.Put("/roles/update/:code", controllers.UpdateRole)
	api.Delete("/roles/delete/:code", controllers.DeleteRole)

	api.Get("/categories", controllers.GetCategories)
	api.Get("/categories/:code", controllers.GetOneCategory)
	api.Post("/categories/create", controllers.CreateCategory)
	api.Put("/categories/update/:code", controllers.UpdateCategory)
	api.Delete("/categories/delete/:code", controllers.DeleteCategory)

	api.Get("/paymentMethod", controllers.GetPaymentMethods)
	api.Get("/paymentMethod/:code", controllers.GetOnePaymentMethod)
	api.Post("/paymentMethod/create", controllers.CreatePaymentMethod)
	api.Put("/paymentMethod/update/:code", controllers.UpdatePaymentMethod)
	api.Delete("/paymentMethod/delete/:code", controllers.DeletePaymentMethod)

	api.Get("/sales", controllers.GetSales)
	api.Get("/sales/:code", controllers.GetOneSale)
	api.Post("/sales/create", controllers.CreateSale)
	api.Put("/sales/update/:code", controllers.UpdateSale)
	api.Delete("/sales/delete/:code", controllers.DeleteSale)

	api.Get("/salesItems", controllers.GetSaleItems)
	api.Get("/salesItems/:code", controllers.GetOneSaleItem)
	api.Post("/salesItems/create", controllers.CreateSaleItem)
	api.Put("/salesItems/update/:code", controllers.UpdateSaleItem)
	api.Delete("/salesItems/delete/:code", controllers.DeleteSaleItem)
}
