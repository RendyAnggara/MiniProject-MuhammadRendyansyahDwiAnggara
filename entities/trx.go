package entities

import "time"

type Transaction struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID            uint      `gorm:"not null" json:"-"`
	ShippingAddressID uint      `gorm:"not null" json:"-"`
	TotalPrice        int       `gorm:"not null" json:"harga_total"`
	InvoiceCode       string    `gorm:"type:varchar(100);unique" json:"kode_invoice"`
	PaymentMethod     string    `gorm:"type:varchar(50)" json:"method_bayar"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`

	ShippingAddress    Address             `gorm:"foreignKey:ShippingAddressID" json:"alamat_kirim"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID" json:"detail_trx"`
}

type TransactionDetail struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	TransactionID uint      `gorm:"not null" json:"-"`
	ProductLogID  uint      `gorm:"not null" json:"-"`
	StoreID       uint      `gorm:"not null" json:"-"`
	Quantity      int       `gorm:"not null" json:"kuantitas"`
	TotalPrice    int       `gorm:"not null" json:"harga_total"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`

	Product ProductLog `gorm:"foreignKey:ProductLogID" json:"product"`
	Store   Store      `gorm:"foreignKey:StoreID" json:"toko"`
}

type ProductLog struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID     uint      `gorm:"not null" json:"-"`
	ProductName   string    `gorm:"type:varchar(255)" json:"nama_produk"`
	Slug          string    `gorm:"type:varchar(255)" json:"slug"`
	ResellerPrice int       `gorm:"not null" json:"harga_reseler"`
	ConsumerPrice int       `gorm:"not null" json:"harga_konsumen"`
	Description   string    `gorm:"type:text" json:"deskripsi"`
	StoreID       uint      `gorm:"not null" json:"-"`
	CategoryID    uint      `gorm:"not null" json:"-"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`

	Store    Store          `gorm:"foreignKey:StoreID" json:"toko"`
	Category Category       `gorm:"foreignKey:CategoryID" json:"category"`
	Photos   []ProductPhoto `gorm:"-" json:"photos"`
}

type CreateTransactionItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"kuantitas"`
}

type CreateTransactionRequest struct {
	PaymentMethod     string                         `json:"method_bayar"`
	ShippingAddressID uint                           `json:"alamat_kirim"`
	DetailTrx         []CreateTransactionItemRequest `json:"detail_trx"`
}