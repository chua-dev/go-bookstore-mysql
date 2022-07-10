package main

import (
	"log"
	"net/http"

	"github.com/chua-dev/go-bookstore-mysql/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// main.go is mainly to create server, and tell and point where router is at

// Create Server
func main() {
	// initialize the router
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	// If there is error say in log
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
