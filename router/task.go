package router

import (
	"github.com/gofiber/fiber/v2"
	"workflow/api"
)

func InitTask(r fiber.Router) {
	r.Post("/add", api.AddTask)
}
