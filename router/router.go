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

	app.Get("/classes", func(c *fiber.Ctx) error {
		data, err := controller.FetchFromSupabase()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Send(data)
	})

	app.Get("/compare", func(c *fiber.Ctx) error {
		text1 := c.Query("text1")
		text2 := c.Query("text2")

		if text1 == "" || text2 == "" {
			return c.Status(fiber.StatusBadRequest).SendString("text1 and text2 query parameters are required")
		}

		score, err := controller.CompareResponse(text1, text2)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"score": score,
		})
	})
}
