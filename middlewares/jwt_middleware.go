package middlewares

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"

	"github.com/gofiber/fiber/v3"
)

func JWTMiddleware(c fiber.Ctx) error {
	tokenString := c.Get("token")
	if tokenString == "" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized", []string{"Token tidak ditemukan"})
	}

	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized", []string{"Token tidak valid atau kadaluarsa"})
	}

	c.Locals("user_id", claims.UserID)
	c.Locals("email", claims.Email)
	c.Locals("is_admin", claims.IsAdmin)

	return c.Next()
}