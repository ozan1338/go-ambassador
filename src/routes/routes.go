package routes

import (
	"go-ambassador/src/controllers"
	"go-ambassador/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("ok")
	})
	api := app.Group("api/ambassador")

	api.Post("register", controllers.Register)
	api.Post("login", controllers.Login)
	api.Get("products/frontend", controllers.ProductsFrontend)
	api.Get("products/backend", controllers.ProductsBackend)

	ambassadorAuthenticated := api.Use(middlewares.IsAuthenticated)
	ambassadorAuthenticated.Get("user", controllers.User)
	ambassadorAuthenticated.Post("logout", controllers.Logout)
	ambassadorAuthenticated.Put("users/info", controllers.UpdateInfo)
	ambassadorAuthenticated.Put("users/password", controllers.UpdatePassword)
	ambassadorAuthenticated.Post("links", controllers.CreateLink)
	ambassadorAuthenticated.Get("stats", controllers.Stats)
	ambassadorAuthenticated.Get("rankings", controllers.Rankings)

}
