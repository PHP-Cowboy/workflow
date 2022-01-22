package router

import "github.com/gofiber/fiber/v2"

func InitCommon(r fiber.Router) {
	r.Get("/header")
}
