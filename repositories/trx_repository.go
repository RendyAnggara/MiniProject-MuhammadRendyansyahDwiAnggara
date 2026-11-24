package repositories

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(userID uint, req entities.CreateTransactionRequest) (uint, error)
	FindAllByUserID(userID uint, search string, limit, offset int) ([]entities.Transaction, error)
	FindByIDAndUserID(id, userID uint) (*entities.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(userID uint, req entities.CreateTransactionRequest) (uint, error) {
	var trxID uint

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Validasi alamat milik user
		var address entities.Address
		if err := tx.Where("id = ? AND user_id = ?", req.ShippingAddressID, userID).First(&address).Error; err != nil {
			return errors.New("alamat_kirim tidak ditemukan atau bukan milik Anda")
		}

		if len(req.DetailTrx) == 0 {
			return errors.New("detail_trx tidak boleh kosong")
		}

		// 2. Buat header transaksi awal
		trx := entities.Transaction{
			UserID:            userID,
			ShippingAddressID: req.ShippingAddressID,
			PaymentMethod:     req.PaymentMethod,
			InvoiceCode:       fmt.Sprintf("INV-%d", time.Now().Unix()),
			TotalPrice:        0,
		}
		if err := tx.Create(&trx).Error; err != nil {
			return err
		}

		grandTotal := 0

		// 3. Proses setiap item produk
		for _, item := range req.DetailTrx {
			if item.Quantity <= 0 {
				return errors.New("kuantitas produk minimal 1")
			}

			var product entities.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				return fmt.Errorf("produk dengan id %d tidak ditemukan", item.ProductID)
			}

			if product.Stock < item.Quantity {
				return fmt.Errorf("stok produk '%s' tidak mencukupi (sisa: %d)", product.ProductName, product.Stock)
			}

			// Kurangi stok produk
			product.Stock -= item.Quantity
			if err := tx.Save(&product).Error; err != nil {
				return err
			}

			// Buat snapshot ke product_logs
			productLog := entities.ProductLog{
				ProductID:     product.ID,
				ProductName:   product.ProductName,
				Slug:          product.Slug,
				ResellerPrice: product.ResellerPrice,
				ConsumerPrice: product.ConsumerPrice,
				Description:   product.Description,
				StoreID:       product.StoreID,
				CategoryID:    product.CategoryID,
			}
			if err := tx.Create(&productLog).Error; err != nil {
				return err
			}

			subTotal := product.ConsumerPrice * item.Quantity
			grandTotal += subTotal

			detail := entities.TransactionDetail{
				TransactionID: trx.ID,
				ProductLogID:  productLog.ID,
				StoreID:       product.StoreID,
				Quantity:      item.Quantity,
				TotalPrice:    subTotal,
			}
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}

		// 4. Update total harga akhir transaksi
		trx.TotalPrice = grandTotal
		if err := tx.Save(&trx).Error; err != nil {
			return err
		}

		trxID = trx.ID
		return nil
	})

	return trxID, err
}

func (r *transactionRepository) populateLogPhotos(transactions []entities.Transaction) {
	for i := range transactions {
		for j := range transactions[i].TransactionDetails {
			prodID := transactions[i].TransactionDetails[j].Product.ProductID
			var photos []entities.ProductPhoto
			r.db.Where("product_id = ?", prodID).Find(&photos)
			transactions[i].TransactionDetails[j].Product.Photos = photos
		}
	}
}

func (r *transactionRepository) FindAllByUserID(userID uint, search string, limit, offset int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	query := r.db.
		Preload("ShippingAddress").
		Preload("TransactionDetails.Store").
		Preload("TransactionDetails.Product.Store").
		Preload("TransactionDetails.Product.Category").
		Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("invoice_code LIKE ?", "%"+search+"%")
	}
	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}

	r.populateLogPhotos(transactions)
	return transactions, nil
}

func (r *transactionRepository) FindByIDAndUserID(id, userID uint) (*entities.Transaction, error) {
	var trx entities.Transaction
	err := r.db.
		Preload("ShippingAddress").
		Preload("TransactionDetails.Store").
		Preload("TransactionDetails.Product.Store").
		Preload("TransactionDetails.Product.Category").
		Where("id = ? AND user_id = ?", id, userID).
		First(&trx).Error
	if err != nil {
		return nil, err
	}

	slice := []entities.Transaction{trx}
	r.populateLogPhotos(slice)
	return &slice[0], nil
}