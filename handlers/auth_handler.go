package handlers

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/services"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req entities.RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}

	if err := h.authService.Register(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to POST data", "Register Succeed")
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req entities.LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}

	res, err := h.authService.Login(req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Failed to POST data", []string{err.Error()})
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to POST data", res)
}