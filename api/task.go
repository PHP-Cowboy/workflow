package api

import "github.com/gofiber/fiber/v2"

func AddTask(c *fiber.Ctx) error {
	return c.SendString("add Task")
}
