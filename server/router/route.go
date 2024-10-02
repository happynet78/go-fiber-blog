package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/happynet78/go-fiber-blog/controller"
)

// Setup Router List
func SetupRoutes(app *fiber.App) {

	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	// list => get
	// add => post
	// update => put
	// delete => delete

	// app.Use(middleware.IsAuthenticate)
	app.Get("/", controller.BlogList)
	app.Get("/blog/", controller.BlogList)
	app.Get("/blog/:id", controller.BlogDetail)
	app.Post("/blog/", controller.BlogCreate)
	app.Put("/blog/:id", controller.BlogUpdate)
	app.Delete("/blog/:id", controller.BlogDelete)
}
