package middlewares

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"

	"github.com/gofiber/fiber/v3"
)

func AdminMiddleware(c fiber.Ctx) error {
	isAdmin, ok := c.Locals("is_admin").(bool)
	if !ok || !isAdmin {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized", []string{"Akses ditolak: Hanya Admin yang diizinkan"})
	}
	return c.Next()
}