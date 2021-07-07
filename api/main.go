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
	return r
}
