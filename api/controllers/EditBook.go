package controllers

import (
	"context"
	"encoding/json"
	"github.com/GoElasticsearch/api/models"
	"github.com/GoElasticsearch/api/utils"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	book := &models.Book{}
	_ = json.NewDecoder(r.Body).Decode(book)

	esClient, err := utils.GetESClient()
	if err != nil {
		log.Fatal("not able to get ES client")
	}

	var b strings.Builder
	b.WriteString(`{"doc": {"title" : "`)
	b.WriteString(book.Title)
	b.WriteString(`",`)
	b.WriteString(`"author" : "`)
	b.WriteString(book.Author)
	b.WriteString(`",`)
	b.WriteString(`"abstract" : "`)
	b.WriteString(book.Abstract)
	b.WriteString(`"}}`)

	req := esapi.UpdateRequest{
		Index:      "books",
		DocumentID: id,
		Body:       strings.NewReader(b.String()),
	}

	do, err := req.Do(context.Background(), esClient)
	if err != nil {
		log.Fatal(err)
	}

	defer do.Body.Close()

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}
