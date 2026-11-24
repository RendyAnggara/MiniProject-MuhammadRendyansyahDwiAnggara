package repositories

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"

	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address *entities.Address) error
	FindAllByUserID(userID uint, titleFilter string) ([]entities.Address, error)
	FindByIDAndUserID(id, userID uint) (*entities.Address, error)
	Update(address *entities.Address) error
	Delete(address *entities.Address) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) Create(address *entities.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepository) FindAllByUserID(userID uint, titleFilter string) ([]entities.Address, error) {
	var addresses []entities.Address
	query := r.db.Where("user_id = ?", userID)
	if titleFilter != "" {
		query = query.Where("title LIKE ?", "%"+titleFilter+"%")
	}
	err := query.Find(&addresses).Error
	return addresses, err
}

func (r *addressRepository) FindByIDAndUserID(id, userID uint) (*entities.Address, error) {
	var address entities.Address
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&address).Error
	return &address, err
}

func (r *addressRepository) Update(address *entities.Address) error {
	return r.db.Save(address).Error
}

func (r *addressRepository) Delete(address *entities.Address) error {
	return r.db.Delete(address).Error
}