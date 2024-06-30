package main

import (
	"log"
	"net/http"

	// "github.com/Sumit/go-bookstore/pkg/routes"
	// "github.com/Sumit/book-store/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sumit/book-store/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
