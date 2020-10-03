package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otiai10/gosseract/v2"

	"github.com/mirusky-dev/ocr-server/handlers"
)

func main() {
	app := fiber.New()

	// Adding a new tesseract client in every request
	app.Use(func(c *fiber.Ctx) error {
		client := gosseract.NewClient()
		c.Context().SetUserValue("tess-client", client)
		return c.Next()
	})

	app.Get("/ping", handlers.Ping)

	app.Post("/ocr", handlers.OCR)

	app.Listen(":3000")
}
