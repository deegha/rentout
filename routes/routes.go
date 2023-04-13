package routes

import (
	"rentoutlkApi/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Get("/user", controllers.User)
	app.Post("/products", controllers.AddProduct)
	app.Put("/products/:id", controllers.UpdateProduct)
	app.Get("/products/:id", controllers.ListAProducts)
	app.Get("/products", controllers.ListProducts)
}
