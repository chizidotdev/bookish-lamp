package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=copia port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Sale{})
	db.Create(&Item{ID: "1", Name: "Perfume 1", BuyingPrice: 100, SellingPrice: 200})

	fmt.Println("Connected to DB")
	DB = db
}
