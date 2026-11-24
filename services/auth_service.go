package services

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"MiniProject-MuhammadRendyansyahDwiAnggara/repositories"
	"MiniProject-MuhammadRendyansyahDwiAnggara/utils"
	"errors"
	"fmt"
	"strings"
)

type AuthService interface {
	Register(req entities.RegisterRequest) error
	Login(req entities.LoginRequest) (*entities.UserProfileResponse, error)
}

type authService struct {
	userRepo        repositories.UserRepository
	provCityService ProvCityService
}

func NewAuthService(userRepo repositories.UserRepository, provCityService ProvCityService) AuthService {
	return &authService{userRepo, provCityService}
}

func (s *authService) Register(req entities.RegisterRequest) error {
	if req.Name == "" || req.Password == "" || req.Phone == "" {
		return errors.New("nama, kata_sandi, dan no_telp wajib diisi")
	}

	if _, err := s.userRepo.FindByPhone(req.Phone); err == nil {
		return fmt.Errorf("no_telp %s sudah terdaftar", req.Phone)
	}
	if req.Email != "" {
		if _, err := s.userRepo.FindByEmail(req.Email); err == nil {
			return fmt.Errorf("email %s sudah terdaftar", req.Email)
		}
	}

	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &entities.User{
		Name:        req.Name,
		Password:    hashedPass,
		Phone:       req.Phone,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		About:       req.About,
		Job:         req.Job,
		Email:       req.Email,
		ProvinceID:  req.ProvinceID,
		CityID:      req.CityID,
		IsAdmin:     false,
	}

	storeName := "toko-" + strings.ToLower(strings.ReplaceAll(req.Name, " ", "-"))
	store := &entities.Store{
		StoreName: storeName,
	}

	return s.userRepo.CreateWithStore(user, store)
}

func (s *authService) Login(req entities.LoginRequest) (*entities.UserProfileResponse, error) {
	user, err := s.userRepo.FindByPhone(req.Phone)
	if err != nil {
		return nil, errors.New("No Telp atau kata sandi salah")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("No Telp atau kata sandi salah")
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return nil, err
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
		Token:       token,
	}, nil
}