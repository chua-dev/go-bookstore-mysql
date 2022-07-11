package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chua-dev/go-bookstore-mysql/pkg/models"
	"github.com/chua-dev/go-bookstore-mysql/pkg/utils"
	"github.com/gorilla/mux"
)

// Creating new book as struct from model
var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	// json.Marshal convert the slice newBooks to json
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) // Write the response with res which the json newBooks
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //Retrieve vars form request body
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) // conver id to int
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Context-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Declare Book Struct Type
	CreateBook := &models.Book{}

	// Receive Json Params from Request
	utils.ParseBody(r, CreateBook)

	// Create Book and retrieve book struct object
	b := CreateBook.CreateBook()

	// Convert the book struct object to json for return
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Get ID from request Params
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing int")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Declare an empty book
	var updateBook = &models.Book{}

	// Parse the Request body params into readable format into empty book above
	utils.ParseBody(r, updateBook)

	// Vars get the ID params from URL
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing ID to int")
	}

	selectedBook, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		selectedBook.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		selectedBook.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		selectedBook.Publication = updateBook.Publication
	}

	db.Save(&selectedBook)
	res, _ := json.Marshal(selectedBook)
	w.Header().Set("Context-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
