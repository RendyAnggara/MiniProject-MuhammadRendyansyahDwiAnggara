package services

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/repositories"
	"errors"
	"regexp"
	"strings"
)

type ProductService interface {
	Create(userID uint, req entities.Product, photoURLs []string) (uint, error)
	GetAll(filter repositories.ProductFilter) ([]entities.Product, error)
	GetByID(id uint) (*entities.Product, error)
	Update(productID, userID uint, req entities.Product, photoURLs []string) error
	Delete(productID, userID uint) error
}

type productService struct {
	productRepo  repositories.ProductRepository
	storeRepo    repositories.StoreRepository
	categoryRepo repositories.CategoryRepository
}

func NewProductService(
	productRepo repositories.ProductRepository,
	storeRepo repositories.StoreRepository,
	categoryRepo repositories.CategoryRepository,
) ProductService {
	return &productService{productRepo, storeRepo, categoryRepo}
}

func generateSlug(text string) string {
	slug := strings.ToLower(strings.TrimSpace(text))
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

func (s *productService) Create(userID uint, req entities.Product, photoURLs []string) (uint, error) {
	store, err := s.storeRepo.FindByUserID(userID)
	if err != nil {
		return 0, errors.New("toko user tidak ditemukan")
	}

	if _, err := s.categoryRepo.FindByID(req.CategoryID); err != nil {
		return 0, errors.New("kategori tidak ditemukan")
	}

	req.StoreID = store.ID
	req.Slug = generateSlug(req.ProductName)

	for _, url := range photoURLs {
		req.Photos = append(req.Photos, entities.ProductPhoto{URL: url})
	}

	if err := s.productRepo.Create(&req); err != nil {
		return 0, err
	}
	return req.ID, nil
}

func (s *productService) GetAll(filter repositories.ProductFilter) ([]entities.Product, error) {
	return s.productRepo.FindAll(filter)
}

func (s *productService) GetByID(id uint) (*entities.Product, error) {
	prod, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("No Data Product")
	}
	return prod, nil
}

func (s *productService) Update(productID, userID uint, req entities.Product, photoURLs []string) error {
	store, err := s.storeRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("toko user tidak ditemukan")
	}

	prod, err := s.productRepo.FindByID(productID)
	if err != nil {
		return errors.New("record not found")
	}

	if prod.StoreID != store.ID {
		return errors.New("Unauthorized: Anda tidak memiliki akses ke produk ini")
	}

	if req.ProductName != "" {
		prod.ProductName = req.ProductName
		prod.Slug = generateSlug(req.ProductName)
	}
	if req.CategoryID > 0 {
		if _, err := s.categoryRepo.FindByID(req.CategoryID); err != nil {
			return errors.New("kategori tidak ditemukan")
		}
		prod.CategoryID = req.CategoryID
	}
	if req.ResellerPrice > 0 {
		prod.ResellerPrice = req.ResellerPrice
	}
	if req.ConsumerPrice > 0 {
		prod.ConsumerPrice = req.ConsumerPrice
	}
	if req.Stock > 0 {
		prod.Stock = req.Stock
	}
	if req.Description != "" {
		prod.Description = req.Description
	}

	var newPhotos []entities.ProductPhoto
	for _, url := range photoURLs {
		newPhotos = append(newPhotos, entities.ProductPhoto{URL: url})
	}

	return s.productRepo.Update(prod, newPhotos)
}

func (s *productService) Delete(productID, userID uint) error {
	store, err := s.storeRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("toko user tidak ditemukan")
	}

	prod, err := s.productRepo.FindByID(productID)
	if err != nil {
		return errors.New("record not found")
	}

	if prod.StoreID != store.ID {
		return errors.New("Unauthorized: Anda tidak memiliki akses ke produk ini")
	}

	return s.productRepo.Delete(prod)
}