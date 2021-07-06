package controllers

import (
	"fmt"
	"github.com/GoElasticsearch/api/utils"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
	"strings"
)


func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	esClient, err := utils.GetESClient()
	if err != nil{
		log.Fatal("not able to get ES client")
	}
	log.Println(elasticsearch.Version)
	log.Println(esClient.Info())
	//var books []books.Book

	res, err := esClient.Index(
		"books",
		strings.NewReader(`{"title":"doe"}`),
		)
	log.Println(res)


	fmt.Fprintf(w, "welcome on the search page blablablu")
	log.Println("SearchBooks")
}
