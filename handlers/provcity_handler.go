package handlers

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/services"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"

	"github.com/gofiber/fiber/v3"
)

type ProvCityHandler struct {
	service services.ProvCityService
}

func NewProvCityHandler(service services.ProvCityService) *ProvCityHandler {
	return &ProvCityHandler{service}
}

func (h *ProvCityHandler) GetListProvinces(c fiber.Ctx) error {
	data, err := h.service.GetListProvinces()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to get data", data)
}

func (h *ProvCityHandler) GetListCities(c fiber.Ctx) error {
	provID := c.Params("prov_id")
	data, err := h.service.GetListCities(provID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to get data", data)
}

func (h *ProvCityHandler) GetDetailProvince(c fiber.Ctx) error {
	provID := c.Params("prov_id")
	data, err := h.service.GetDetailProvince(provID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to get data", data)
}

func (h *ProvCityHandler) GetDetailCity(c fiber.Ctx) error {
	cityID := c.Params("city_id")
	data, err := h.service.GetDetailCity(cityID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to get data", data)
}