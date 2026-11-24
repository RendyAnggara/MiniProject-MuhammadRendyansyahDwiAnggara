package handlers

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/services"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service}
}

func (h *TransactionHandler) GetAll(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	search := c.Query("search")
	page := fiber.Query[int](c, "page", 0)
	limit := fiber.Query[int](c, "limit", 0)

	data, err := h.service.GetMyTransactions(userID, search, page, limit)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}

	res := fiber.Map{
		"data":  data,
		"page":  page,
		"limit": limit,
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", res)
}

func (h *TransactionHandler) GetByID(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	data, err := h.service.GetByID(uint(id), userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", data)
}

func (h *TransactionHandler) Create(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var req entities.CreateTransactionRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}

	id, err := h.service.Create(userID, req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to POST data", id)
}