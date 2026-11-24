package handlers

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/repositories"
	"MiniProject-MuhammadRendyansyahDwiAnggara/services"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) GetAll(c fiber.Ctx) error {
	page := fiber.Query[int](c, "page", 1)
	limit := fiber.Query[int](c, "limit", 10)
	offset := (page - 1) * limit

	filter := repositories.ProductFilter{
		ProductName: c.Query("nama_produk"),
		CategoryID:  fiber.Query[int](c, "category_id", 0),
		StoreID:     fiber.Query[int](c, "toko_id", 0),
		MinPrice:    fiber.Query[int](c, "min_harga", 0),
		MaxPrice:    fiber.Query[int](c, "max_harga", 0),
		Limit:       limit,
		Offset:      offset,
	}

	products, err := h.service.GetAll(filter)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}

	res := fiber.Map{
		"data":  products,
		"page":  page,
		"limit": limit,
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", res)
}

func (h *ProductHandler) GetByID(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	prod, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", prod)
}

func (h *ProductHandler) Create(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	catID, _ := strconv.Atoi(c.FormValue("category_id"))
	resellerPrice, _ := strconv.Atoi(c.FormValue("harga_reseller"))
	consumerPrice, _ := strconv.Atoi(c.FormValue("harga_konsumen"))
	stock, _ := strconv.Atoi(c.FormValue("stok"))

	req := entities.Product{
		ProductName:   c.FormValue("nama_produk"),
		CategoryID:    uint(catID),
		ResellerPrice: resellerPrice,
		ConsumerPrice: consumerPrice,
		Stock:         stock,
		Description:   c.FormValue("deskripsi"),
	}

	var photoURLs []string
	if form, err := c.MultipartForm(); err == nil && form != nil {
		files := form.File["photos"]
		for _, file := range files {
			savedName, err := utils.SaveUploadedFile(c, file)
			if err != nil {
				return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
			}
			photoURLs = append(photoURLs, savedName)
		}
	}

	id, err := h.service.Create(userID, req, photoURLs)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to POST data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to POST data", id)
}

func (h *ProductHandler) Update(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	catID, _ := strconv.Atoi(c.FormValue("category_id"))
	resellerPrice, _ := strconv.Atoi(c.FormValue("harga_reseller"))
	consumerPrice, _ := strconv.Atoi(c.FormValue("harga_konsumen"))
	stock, _ := strconv.Atoi(c.FormValue("stok"))

	req := entities.Product{
		ProductName:   c.FormValue("nama_produk"),
		CategoryID:    uint(catID),
		ResellerPrice: resellerPrice,
		ConsumerPrice: consumerPrice,
		Stock:         stock,
		Description:   c.FormValue("deskripsi"),
	}

	var photoURLs []string
	if form, err := c.MultipartForm(); err == nil && form != nil {
		files := form.File["photos"]
		for _, file := range files {
			savedName, err := utils.SaveUploadedFile(c, file)
			if err != nil {
				return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to PUT data", []string{err.Error()})
			}
			photoURLs = append(photoURLs, savedName)
		}
	}

	if err := h.service.Update(uint(id), userID, req, photoURLs); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", "")
}

func (h *ProductHandler) Delete(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.service.Delete(uint(id), userID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to GET data", []string{err.Error()})
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "Succeed to GET data", "")
}