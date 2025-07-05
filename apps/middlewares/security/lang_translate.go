package security

import "github.com/gofiber/fiber/v2"

func (uc *security) LangTranslate(c *fiber.Ctx) error {

	lang := c.Query("lang")
	if lang == "" {
		lang = "en"
	}

	c.Locals("lang", lang)

	return c.Next()
}
