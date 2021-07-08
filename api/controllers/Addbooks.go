package controllers

import (
	"context"
	"github.com/GoElasticsearch/api/utils"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"net/http"
	"strings"
)

func AddBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Add in elasticsearch")
	esClient, err := utils.GetESClient()
	if err != nil {
		log.Fatalf("Error getting EsClient: %v", err)
	}
	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         "books",        // The default index name
		Client:        esClient,       // The Elasticsearch client
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
	if err != nil{
		log.Fatalf("Error bulk add: %v", err)
	}
	indexer.Close(context.Background())


}
