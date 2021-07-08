package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/GoElasticsearch/api/utils"
	"log"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")

	esClient, err := utils.GetESClient()
	if err != nil {
		log.Fatal("not able to get ES client")
	}

	var buf bytes.Buffer

	//create the query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": keyword,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	response, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex("books"),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer response.Body.Close()

	if response.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				response.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var books map[string]interface{}

	res, _ := esClient.Search(esClient.Search.WithTrackTotalHits(true))
	if err := json.NewDecoder(res.Body).Decode(&books); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response.String()))
}
