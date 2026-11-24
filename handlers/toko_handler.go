package handlers

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/services"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type StoreHandler struct {
	service services.StoreService
}

func NewStoreHandler(service services.StoreService) *StoreHandler {
	return &StoreHandler{service}
}

func (h *StoreHandler) GetMyStore(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	data, err := h.service.GetMyStore(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", data)
}

func (h *StoreHandler) GetByID(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id_toko"))
	data, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Failed to GET data", []string{err.Error()})
	}

	// Response Get Toko By ID sesuai Postman tidak menyertakan user_id
	res := fiber.Map{
		"id":        data.ID,
		"nama_toko": data.StoreName,
		"url_foto":  data.PhotoURL,
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", res)
}

func (h *StoreHandler) GetAll(c fiber.Ctx) error {
	page := fiber.Query[int](c, "page", 1)
	limit := fiber.Query[int](c, "limit", 10)
	nameFilter := c.Query("nama")

	stores, err := h.service.GetAll(page, limit, nameFilter)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}

	var list []fiber.Map
	for _, s := range stores {
		list = append(list, fiber.Map{
			"id":        s.ID,
			"nama_toko": s.StoreName,
			"url_foto":  s.PhotoURL,
		})
	}

	res := fiber.Map{
		"page":  page,
		"limit": limit,
		"data":  list,
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", res)
}

func (h *StoreHandler) UpdateStore(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	storeID, _ := strconv.Atoi(c.Params("id_toko"))
	storeName := c.FormValue("nama_toko")

	var photoURL string
	file, err := c.FormFile("photo")
	if err == nil && file != nil {
		savedName, err := utils.SaveUploadedFile(c, file)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to UPDATE data", []string{err.Error()})
		}
		photoURL = savedName
	}

	if err := h.service.UpdateStore(uint(storeID), userID, storeName, photoURL); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to UPDATE data", []string{err.Error()})
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to UPDATE data", "Update toko succeed")
}