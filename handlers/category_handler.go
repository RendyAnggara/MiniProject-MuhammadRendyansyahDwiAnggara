package handlers

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/services"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) GetAll(c fiber.Ctx) error {
	data, err := h.service.GetAll()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", data)
}

func (h *CategoryHandler) GetByID(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", data)
}

func (h *CategoryHandler) Create(c fiber.Ctx) error {
	var req entities.Category
	if err := c.Bind().Body(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}

	id, err := h.service.Create(req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to POST data", id)
}

func (h *CategoryHandler) Update(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req entities.Category
	if err := c.Bind().Body(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to PUT data", []string{err.Error()})
	}

	if err := h.service.Update(uint(id), req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", "")
}

func (h *CategoryHandler) Delete(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", "")
}