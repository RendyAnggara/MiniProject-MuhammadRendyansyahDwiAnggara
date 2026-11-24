package config

import (
	"fmt"
	"log"
	"os"

	"MiniProject-MuhammadRendyansyahDwiAnggara/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database MySQL: ", err)
	}

	DB = db
	log.Println("Berhasil terhubung ke database MySQL!")
}

func MigrateDatabase() {
	err := DB.AutoMigrate(
		&entities.User{},
		&entities.Store{},
		&entities.Address{},
		&entities.Category{},
		&entities.Product{},
		&entities.ProductPhoto{},
		&entities.Transaction{},
		&entities.TransactionDetail{},
		&entities.ProductLog{},
	)
	if err != nil {
		log.Fatal("Gagal migrasi database: ", err)
	}
	log.Println("Migrasi tabel database berhasil!")
}