package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"copia/api/apps/api/pkg/models"
)

func Init() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	log.Println("Connected to DB")

	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Sale{})
	// db.Create(&models.Item{ID: "1", Name: "Perfume 1", BuyingPrice: 100, SellingPrice: 200})

	return db
}
