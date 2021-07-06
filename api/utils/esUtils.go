package utils

import "github.com/elastic/go-elasticsearch/v8"

func GetESClient() (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewDefaultClient()
	return es, err
}

