package routes

import (
	"gofiber-mongodb/controllers"

	"github.com/gofiber/fiber/v2"
)

func BahanRoute(app *fiber.App) {
	app.Get("/bahan", controllers.GetAllBahan)
}
