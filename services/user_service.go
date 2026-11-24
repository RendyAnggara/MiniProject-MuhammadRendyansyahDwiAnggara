package services

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/repositories"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"
	"errors"
)

type UserService interface {
	GetMyProfile(userID uint) (*entities.UserProfileResponse, error)
	UpdateProfile(userID uint, req entities.RegisterRequest) error
}

type userService struct {
	userRepo        repositories.UserRepository
	provCityService ProvCityService
}

func NewUserService(userRepo repositories.UserRepository, provCityService ProvCityService) UserService {
	return &userService{userRepo, provCityService}
}

func (s *userService) GetMyProfile(userID uint) (*entities.UserProfileResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	prov, _ := s.provCityService.GetDetailProvince(user.ProvinceID)
	city, _ := s.provCityService.GetDetailCity(user.CityID)

	return &entities.UserProfileResponse{
		Name:        user.Name,
		Phone:       user.Phone,
		DateOfBirth: user.DateOfBirth,
		About:       user.About,
		Job:         user.Job,
		Email:       user.Email,
		ProvinceID:  prov,
		CityID:      city,
	}, nil
}

func (s *userService) UpdateProfile(userID uint, req entities.RegisterRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	if req.Phone != "" && req.Phone != user.Phone {
		if existing, err := s.userRepo.FindByPhone(req.Phone); err == nil && existing.ID != userID {
			return errors.New("no_telp sudah digunakan user lain")
		}
		user.Phone = req.Phone
	}

	if req.Email != "" && req.Email != user.Email {
		if existing, err := s.userRepo.FindByEmail(req.Email); err == nil && existing.ID != userID {
			return errors.New("email sudah digunakan user lain")
		}
		user.Email = req.Email
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Password != "" {
		hashed, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user.Password = hashed
	}
	if req.DateOfBirth != "" {
		user.DateOfBirth = req.DateOfBirth
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}
	if req.About != "" {
		user.About = req.About
	}
	if req.Job != "" {
		user.Job = req.Job
	}
	if req.ProvinceID != "" {
		user.ProvinceID = req.ProvinceID
	}
	if req.CityID != "" {
		user.CityID = req.CityID
	}

	return s.userRepo.Update(user)
}