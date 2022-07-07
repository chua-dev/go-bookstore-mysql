package models

import (
	"github.com/chua-dev/go-bookstore-mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB // Declare variable called db, type is gorm DB type

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
