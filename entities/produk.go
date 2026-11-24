package entities

import "time"

type Product struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductName   string    `gorm:"type:varchar(255)" json:"nama_produk"`
	Slug          string    `gorm:"type:varchar(255)" json:"slug"`
	ResellerPrice int       `gorm:"not null" json:"harga_reseler"`
	ConsumerPrice int       `gorm:"not null" json:"harga_konsumen"`
	Stock         int       `gorm:"not null" json:"stok"`
	Description   string    `gorm:"type:text" json:"deskripsi"`
	StoreID       uint      `gorm:"not null" json:"-"`
	CategoryID    uint      `gorm:"not null" json:"-"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`

	Store    Store          `gorm:"foreignKey:StoreID" json:"toko"`
	Category Category       `gorm:"foreignKey:CategoryID" json:"category"`
	Photos   []ProductPhoto `gorm:"foreignKey:ProductID" json:"photos"`
}

type ProductPhoto struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	URL       string    `gorm:"type:varchar(255)" json:"url"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}