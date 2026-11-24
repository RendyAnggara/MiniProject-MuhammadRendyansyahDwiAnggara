package entities

import "time"

type Category struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryName string    `gorm:"type:varchar(255)" json:"nama_category"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}