package services

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/repositories"
	"errors"
)

type CategoryService interface {
	Create(req entities.Category) (uint, error)
	GetAll() ([]entities.Category, error)
	GetByID(id uint) (*entities.Category, error)
	Update(id uint, req entities.Category) error
	Delete(id uint) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) Create(req entities.Category) (uint, error) {
	if req.CategoryName == "" {
		return 0, errors.New("nama_category wajib diisi")
	}
	if err := s.repo.Create(&req); err != nil {
		return 0, err
	}
	return req.ID, nil
}

func (s *categoryService) GetAll() ([]entities.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetByID(id uint) (*entities.Category, error) {
	cat, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("No Data Category")
	}
	return cat, nil
}

func (s *categoryService) Update(id uint, req entities.Category) error {
	cat, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("record not found")
	}
	if req.CategoryName != "" {
		cat.CategoryName = req.CategoryName
	}
	return s.repo.Update(cat)
}

func (s *categoryService) Delete(id uint) error {
	cat, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("record not found")
	}
	return s.repo.Delete(cat)
}