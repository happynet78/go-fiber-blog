package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/happynet78/go-fiber-blog/database"
	"github.com/happynet78/go-fiber-blog/router"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file")
	}

	database.ConnectDB()
}

func main() {
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection")
	}

	defer sqlDb.Close()

	port := os.Getenv("PORT")

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})

	app.Static("/static", "./static")

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		// AllowCredentials: true,
	}))

	// app.Get("/", func(c *fiber.Ctx) error {
	//     return c.JSON(fiber.Map{
	//         "message": "Welcome to first web Application",
	//     })
	//     // return c.SendString("Hello World!")
	// })

	router.SetupRoutes(app)

	app.Listen(":" + port)
}
