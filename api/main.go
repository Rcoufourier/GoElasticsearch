package main

import (
	"github.com/GoElasticsearch/api/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	router := initRouter()
	esClient, err := getESClient()
	if err != nil {
		log.Fatal("es client Failed", err)
	}
	log.Println(esClient)
	log.Println("Running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initRouter () *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/search", controllers.Search).Methods("POST")
	return r
}

func getESClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	return es, err
}


