package router

import "github.com/gofiber/fiber/v2"

func InitApproval(r fiber.Router) {
	r.Post("/approval")
}
