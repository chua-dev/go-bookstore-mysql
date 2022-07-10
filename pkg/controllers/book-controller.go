package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chua-dev/go-bookstore-mysql/pkg/models"
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
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
