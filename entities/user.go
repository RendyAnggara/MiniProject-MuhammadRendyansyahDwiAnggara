package entities

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"type:varchar(255)" json:"nama"`
	Password     string    `gorm:"type:varchar(255)" json:"-"`
	Phone        string    `gorm:"type:varchar(20);unique" json:"no_telp"`
	DateOfBirth  string    `gorm:"type:varchar(50)" json:"tanggal_Lahir"`
	Gender       string    `gorm:"type:varchar(20)" json:"jenis_kelamin"`
	About        string    `gorm:"type:text" json:"tentang"`
	Job          string    `gorm:"type:varchar(100)" json:"pekerjaan"`
	Email        string    `gorm:"type:varchar(255);unique" json:"email"`
	ProvinceID   string    `gorm:"type:varchar(10)" json:"id_provinsi"`
	CityID       string    `gorm:"type:varchar(10)" json:"id_kota"`
	IsAdmin      bool      `gorm:"default:false" json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Store     *Store    `gorm:"foreignKey:UserID" json:"toko,omitempty"`
	Addresses []Address `gorm:"foreignKey:UserID" json:"alamat,omitempty"`
}

type RegisterRequest struct {
	Name        string `json:"nama"`
	Password    string `json:"kata_sandi"`
	Phone       string `json:"no_telp"`
	DateOfBirth string `json:"tanggal_Lahir"`
	Gender      string `json:"jenis_kelamin"`
	About       string `json:"tentang"`
	Job         string `json:"pekerjaan"`
	Email       string `json:"email"`
	ProvinceID  string `json:"id_provinsi"`
	CityID      string `json:"id_kota"`
}

type LoginRequest struct {
	Phone    string `json:"no_telp"`
	Password string `json:"kata_sandi"`
}

type ProvinceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CityResponse struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

type UserProfileResponse struct {
	Name        string            `json:"nama"`
	Phone       string            `json:"no_telp"`
	DateOfBirth string            `json:"tanggal_Lahir"`
	About       string            `json:"tentang"`
	Job         string            `json:"pekerjaan"`
	Email       string            `json:"email"`
	ProvinceID  *ProvinceResponse `json:"id_provinsi"`
	CityID      *CityResponse     `json:"id_kota"`
	Token       string            `json:"token,omitempty"`
}