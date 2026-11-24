package entities

import "time"

type Address struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint      `gorm:"not null" json:"-"`
	Title         string    `gorm:"type:varchar(255)" json:"judul_alamat"`
	RecipientName string    `gorm:"type:varchar(255)" json:"nama_penerima"`
	Phone         string    `gorm:"type:varchar(20)" json:"no_telp"`
	DetailAddress string    `gorm:"type:text" json:"detail_alamat"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}