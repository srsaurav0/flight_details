package dao

import (
	"log"

	"github.com/beego/beego/v2/server/web"
	"github.com/elastic/go-elasticsearch/v8"
)

func GetElasticClient() *elasticsearch.Client {
	ES_API_KEY, _ := web.AppConfig.String("ES_LOCAL_API_KEY")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		APIKey: ES_API_KEY,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}
	return es
}
