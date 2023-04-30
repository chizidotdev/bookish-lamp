package handlers

import "gorm.io/gorm"

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	return handler{db}
}
