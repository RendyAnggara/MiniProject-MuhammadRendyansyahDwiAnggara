package repositories

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"

	"gorm.io/gorm"
)

type StoreRepository interface {
	FindByUserID(userID uint) (*entities.Store, error)
	FindByID(id uint) (*entities.Store, error)
	FindAll(limit, offset int, nameFilter string) ([]entities.Store, error)
	Update(store *entities.Store) error
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{db}
}

func (r *storeRepository) FindByUserID(userID uint) (*entities.Store, error) {
	var store entities.Store
	err := r.db.Where("user_id = ?", userID).First(&store).Error
	return &store, err
}

func (r *storeRepository) FindByID(id uint) (*entities.Store, error) {
	var store entities.Store
	err := r.db.First(&store, id).Error
	return &store, err
}

func (r *storeRepository) FindAll(limit, offset int, nameFilter string) ([]entities.Store, error) {
	var stores []entities.Store
	query := r.db.Limit(limit).Offset(offset)
	if nameFilter != "" {
		query = query.Where("store_name LIKE ?", "%"+nameFilter+"%")
	}
	err := query.Find(&stores).Error
	return stores, err
}

func (r *storeRepository) Update(store *entities.Store) error {
	return r.db.Save(store).Error
}