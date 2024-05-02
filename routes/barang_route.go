package routes

import (
	"gofiber-mongodb/controllers"

	"github.com/gofiber/fiber/v2"
)

func BarangRoute(app *fiber.App) {
	app.Post("/barang", controllers.CreateBarang)
	app.Get("/barang", controllers.GetAllBarang)
}
