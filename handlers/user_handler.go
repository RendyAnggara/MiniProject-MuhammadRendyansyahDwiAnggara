package handlers

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/services"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetMyProfile(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	res, err := h.userService.GetMyProfile(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", res)
}

func (h *UserHandler) UpdateProfile(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var req entities.RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to PUT data", []string{err.Error()})
	}

	if err := h.userService.UpdateProfile(userID, req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to PUT data", []string{err.Error()})
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to PUT data", "")
}