package handlers

import "github.com/gofiber/fiber/v2"

// Ping handle a ping-pong message
func Ping(c *fiber.Ctx) error {
	return c.SendString("Pong!")
}
