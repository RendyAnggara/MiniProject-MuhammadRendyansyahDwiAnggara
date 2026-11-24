package services

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProvCityService interface {
	GetListProvinces() ([]entities.ProvinceResponse, error)
	GetListCities(provID string) ([]entities.CityResponse, error)
	GetDetailProvince(provID string) (*entities.ProvinceResponse, error)
	GetDetailCity(cityID string) (*entities.CityResponse, error)
}

type provCityService struct {
	baseURL string
}

func NewProvCityService() ProvCityService {
	return &provCityService{
		baseURL: "https://www.emsifa.com/api-wilayah-indonesia/api",
	}
}

func (s *provCityService) GetListProvinces() ([]entities.ProvinceResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/provinces.json", s.baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []entities.ProvinceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *provCityService) GetListCities(provID string) ([]entities.CityResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/regencies/%s.json", s.baseURL, provID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []entities.CityResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *provCityService) GetDetailProvince(provID string) (*entities.ProvinceResponse, error) {
	if provID == "" {
		return nil, nil
	}
	resp, err := http.Get(fmt.Sprintf("%s/province/%s.json", s.baseURL, provID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result entities.ProvinceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *provCityService) GetDetailCity(cityID string) (*entities.CityResponse, error) {
	if cityID == "" {
		return nil, nil
	}
	resp, err := http.Get(fmt.Sprintf("%s/regency/%s.json", s.baseURL, cityID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result entities.CityResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}