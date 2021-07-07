package utils

import "github.com/elastic/go-elasticsearch/v8"

func GetESClient() (*elasticsearch.Client, error) {

	config := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
		},
	}

	es, err := elasticsearch.NewClient(config)
	return es, err
}
