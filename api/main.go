package main

import (
	"fmt"
	"github.com/GoElasticsearch/api/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := initRouter()
	log.Println("Running on localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}

}

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/search", controllers.Search).Methods("GET")
	r.HandleFunc("/addfromapi", controllers.AddDataAPI).Methods("GET")
	return r
}
