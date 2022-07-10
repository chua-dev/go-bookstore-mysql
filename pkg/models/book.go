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

// Initialize DB to talk to database
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// gorm has newrecord function
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// Return a list(slice) of Book Struct
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
