package controllers

import (
	"context"
	"github.com/GoElasticsearch/api/utils"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	esClient, err := utils.GetESClient()
	if err != nil {
		log.Fatal("not able to get ES client")
	}

	req := esapi.DeleteRequest{
		Index:      "books",
		DocumentID: id,
	}

	do, err := req.Do(context.Background(), esClient)
	if err != nil {
		log.Fatal(err)
	}
	defer do.Body.Close()

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}
