package services

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/repositories"
	"errors"
)

type StoreService interface {
	GetMyStore(userID uint) (*entities.Store, error)
	GetByID(id uint) (*entities.Store, error)
	GetAll(page, limit int, nameFilter string) ([]entities.Store, error)
	UpdateStore(storeID, userID uint, storeName, photoURL string) error
}

type storeService struct {
	repo repositories.StoreRepository
}

func NewStoreService(repo repositories.StoreRepository) StoreService {
	return &storeService{repo}
}

func (s *storeService) GetMyStore(userID uint) (*entities.Store, error) {
	store, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("Toko tidak ditemukan")
	}
	return store, nil
}

func (s *storeService) GetByID(id uint) (*entities.Store, error) {
	store, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("Toko tidak ditemukan")
	}
	return store, nil
}

func (s *storeService) GetAll(page, limit int, nameFilter string) ([]entities.Store, error) {
	offset := (page - 1) * limit
	return s.repo.FindAll(limit, offset, nameFilter)
}

func (s *storeService) UpdateStore(storeID, userID uint, storeName, photoURL string) error {
	store, err := s.repo.FindByID(storeID)
	if err != nil {
		return errors.New("Toko tidak ditemukan")
	}
	if store.UserID != userID {
		return errors.New("Unauthorized: Anda tidak memiliki akses ke toko ini")
	}

	if storeName != "" {
		store.StoreName = storeName
	}
	if photoURL != "" {
		store.PhotoURL = photoURL
	}

	return s.repo.Update(store)
}