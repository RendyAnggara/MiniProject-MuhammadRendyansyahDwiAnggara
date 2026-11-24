package entities

import "time"

type Store struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	StoreName string    `gorm:"type:varchar(255)" json:"nama_toko"`
	PhotoURL  string    `gorm:"type:varchar(255)" json:"url_foto"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}