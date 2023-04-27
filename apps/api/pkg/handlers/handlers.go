package handlers

import "gorm.io/gorm"

type handler struct {
	DB *gorm.DB
}

func new(db *gorm.DB) handler {
	return handler{db}
}
