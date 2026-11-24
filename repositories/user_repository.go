package repositories

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateWithStore(user *entities.User, store *entities.Store) error
	FindByPhone(phone string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByID(id uint) (*entities.User, error)
	Update(user *entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateWithStore(user *entities.User, store *entities.Store) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		store.UserID = user.ID
		if err := tx.Create(store).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *userRepository) FindByPhone(phone string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByID(id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(user *entities.User) error {
	return r.db.Save(user).Error
}