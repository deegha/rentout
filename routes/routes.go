package routes

import (
	"rentoutlkApi/controllers/product"
	"rentoutlkApi/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", user.Register)
	app.Post("/user", user.CreateUser)
	app.Post("/login", user.Login)
	app.Post("/logout", user.Logout)
	app.Get("/user", user.User)
	app.Post("/products", product.AddProduct)
	app.Put("/products/:id", product.UpdateProduct)
	app.Get("/products/:id", product.ListAProducts)
	app.Get("/products", product.ListProducts)
	app.Get("/products/:id/images", product.ListImages)
}
