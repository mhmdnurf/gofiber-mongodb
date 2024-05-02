package main

import (
	"gofiber-mongodb/configs"
	"gofiber-mongodb/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configs.ConnectDB()
	routes.BarangRoute(app)
	app.Listen(":3000")
}
