package services

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/repositories"
	"errors"
)

type AddressService interface {
	Create(userID uint, req entities.Address) (uint, error)
	GetMyAddresses(userID uint, titleFilter string) ([]entities.Address, error)
	GetByID(id, userID uint) (*entities.Address, error)
	Update(id, userID uint, req entities.Address) error
	Delete(id, userID uint) error
}

type addressService struct {
	repo repositories.AddressRepository
}

func NewAddressService(repo repositories.AddressRepository) AddressService {
	return &addressService{repo}
}

func (s *addressService) Create(userID uint, req entities.Address) (uint, error) {
	req.UserID = userID
	if err := s.repo.Create(&req); err != nil {
		return 0, err
	}
	return req.ID, nil
}

func (s *addressService) GetMyAddresses(userID uint, titleFilter string) ([]entities.Address, error) {
	return s.repo.FindAllByUserID(userID, titleFilter)
}

func (s *addressService) GetByID(id, userID uint) (*entities.Address, error) {
	addr, err := s.repo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, errors.New("record not found")
	}
	return addr, nil
}

func (s *addressService) Update(id, userID uint, req entities.Address) error {
	addr, err := s.repo.FindByIDAndUserID(id, userID)
	if err != nil {
		return errors.New("record not found")
	}

	if req.Title != "" {
		addr.Title = req.Title
	}
	if req.RecipientName != "" {
		addr.RecipientName = req.RecipientName
	}
	if req.Phone != "" {
		addr.Phone = req.Phone
	}
	if req.DetailAddress != "" {
		addr.DetailAddress = req.DetailAddress
	}

	return s.repo.Update(addr)
}

func (s *addressService) Delete(id, userID uint) error {
	addr, err := s.repo.FindByIDAndUserID(id, userID)
	if err != nil {
		return errors.New("record not found")
	}
	return s.repo.Delete(addr)
}