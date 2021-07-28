package main

import (
	"github.com/GoElasticsearch/api/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := initRouter()
	log.Println("Running on localhost:8080")
	http.ListenAndServe(":8080", router)
}

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/search", controllers.Search).Methods("GET")
	r.HandleFunc("/all-books", controllers.SearchAll).Methods("GET")
	r.HandleFunc("/add-book", controllers.AddBook).Methods("POST")
	r.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")
	r.HandleFunc("/edit/{id}", controllers.Edit).Methods("PUT")
	r.HandleFunc("/add-from-api", controllers.AddDataAPI).Methods("GET")
	return r
}
