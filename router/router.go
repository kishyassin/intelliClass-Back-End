package router

import (
	"intelliClass/controller"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	//health
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ENSET DREAMERS TO THE MOON")
	})
	app.Get("/profiles", func(c *fiber.Ctx) error {
		data, err := controller.FetchFromSupabase()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Send(data)
	})
}
