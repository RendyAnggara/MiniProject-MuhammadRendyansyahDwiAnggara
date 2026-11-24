package handlers

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/services"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type AddressHandler struct {
	service services.AddressService
}

func NewAddressHandler(service services.AddressService) *AddressHandler {
	return &AddressHandler{service}
}

func (h *AddressHandler) GetMyAddresses(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	titleFilter := c.Query("judul_alamat")

	data, err := h.service.GetMyAddresses(userID, titleFilter)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", data)
}

func (h *AddressHandler) GetByID(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	data, err := h.service.GetByID(uint(id), userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", data)
}

func (h *AddressHandler) Create(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var req entities.Address
	if err := c.Bind().Body(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}

	id, err := h.service.Create(userID, req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to POST data", id)
}

func (h *AddressHandler) Update(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	var req entities.Address
	if err := c.Bind().Body(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to PUT data", []string{err.Error()})
	}

	if err := h.service.Update(uint(id), userID, req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", "")
}

func (h *AddressHandler) Delete(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.service.Delete(uint(id), userID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", "")
}