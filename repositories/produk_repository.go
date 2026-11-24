package repositories

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"

	"gorm.io/gorm"
)

type ProductFilter struct {
	ProductName string
	CategoryID  int
	StoreID     int
	MinPrice    int
	MaxPrice    int
	Limit       int
	Offset      int
}

type ProductRepository interface {
	Create(product *entities.Product) error
	FindAll(filter ProductFilter) ([]entities.Product, error)
	FindByID(id uint) (*entities.Product, error)
	Update(product *entities.Product, newPhotos []entities.ProductPhoto) error
	Delete(product *entities.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *entities.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll(filter ProductFilter) ([]entities.Product, error) {
	var products []entities.Product
	query := r.db.Preload("Store").Preload("Category").Preload("Photos")

	if filter.ProductName != "" {
		query = query.Where("product_name LIKE ?", "%"+filter.ProductName+"%")
	}
	if filter.CategoryID > 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.StoreID > 0 {
		query = query.Where("store_id = ?", filter.StoreID)
	}
	if filter.MinPrice > 0 {
		query = query.Where("consumer_price >= ?", filter.MinPrice)
	}
	if filter.MaxPrice > 0 {
		query = query.Where("consumer_price <= ?", filter.MaxPrice)
	}

	err := query.Limit(filter.Limit).Offset(filter.Offset).Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*entities.Product, error) {
	var product entities.Product
	err := r.db.Preload("Store").Preload("Category").Preload("Photos").First(&product, id).Error
	return &product, err
}

func (r *productRepository) Update(product *entities.Product, newPhotos []entities.ProductPhoto) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(product).Error; err != nil {
			return err
		}
		if len(newPhotos) > 0 {
			if err := tx.Where("product_id = ?", product.ID).Delete(&entities.ProductPhoto{}).Error; err != nil {
				return err
			}
			for i := range newPhotos {
				newPhotos[i].ProductID = product.ID
			}
			if err := tx.Create(&newPhotos).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *productRepository) Delete(product *entities.Product) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("product_id = ?", product.ID).Delete(&entities.ProductPhoto{}).Error; err != nil {
			return err
		}
		return tx.Delete(product).Error
	})
}