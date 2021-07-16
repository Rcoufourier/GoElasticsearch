package controllers

import (
	"context"
	"encoding/json"
	"github.com/GoElasticsearch/api/models"
	"github.com/GoElasticsearch/api/utils"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net/http"
	"strings"
)

func AddOneBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	_ = json.NewDecoder(r.Body).Decode(book)

	esClient, err := utils.GetESClient()
	if err != nil {
		log.Fatal("not able to get ES client")
	}

	var b strings.Builder
	b.WriteString(`{"title" : "`)
	b.WriteString(book.Title)
	b.WriteString(`",`)
	b.WriteString(`"author" : "`)
	b.WriteString(book.Author)
	b.WriteString(`",`)
	b.WriteString(`"abstract" : "`)
	b.WriteString(book.Abstract)
	b.WriteString(`"}`)

	req := esapi.IndexRequest{
		Index: "books",
		Body:  strings.NewReader(b.String()),
	}

	do, err := req.Do(context.Background(), esClient)
	if err != nil {
		log.Fatal(err)
	}
	defer do.Body.Close()

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}
