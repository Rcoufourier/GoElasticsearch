package controllers

import (
	"context"
	"encoding/json"
	"github.com/GoElasticsearch/api/models"
	"github.com/GoElasticsearch/api/utils"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"net/http"
	"strings"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
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

func AddBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Add in elasticsearch")
	esClient, err := utils.GetESClient()
	if err != nil {
		log.Fatalf("Error getting EsClient: %v", err)
	}
	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  "books",  // The default index name
		Client: esClient, // The Elasticsearch client
	})
	if err != nil {
		log.Fatalf("Error bulk indexer: %v", err)
	}
	err = indexer.Add(
		context.Background(),
		esutil.BulkIndexerItem{
			Action: "index",
			Body:   strings.NewReader(`{ "title":"Michel doe","author":"JeanBonneau", "abstract":"Lorem2 blablou doe ipsum dolor sit amet, consectetur adipiscing elit." }`),
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
				log.Printf("item: %s", item)
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
				if err != nil {
					log.Printf("ERROR: %s", err)
				} else {
					log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
				}
			},
		})
	if err != nil {
		log.Fatalf("Error bulk add: %v", err)
	}
	indexer.Close(context.Background())
}
