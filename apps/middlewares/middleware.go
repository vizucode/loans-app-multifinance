package middlewares

import "github.com/gofiber/fiber/v2"

type IMiddleware interface {
	AuthMiddleware(c *fiber.Ctx) error
	LangTranslate(c *fiber.Ctx) error
}
