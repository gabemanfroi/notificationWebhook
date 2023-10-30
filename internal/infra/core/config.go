package core

import (
	constants "github.com/gabemanfroi/notificationWebhook/internal/infra"
	"log"
	"os"
)

type Config struct {
	Elasticsearch struct {
		Host     string
		Username string
		Password string
	}
}

var AppConfig = Config{}

func LoadConfig() {
	log.Println("Loading config...")
	AppConfig.Elasticsearch.Host = os.Getenv(constants.ElasticsearchHost)
	AppConfig.Elasticsearch.Username = os.Getenv(constants.ElasticsearchUsername)
	AppConfig.Elasticsearch.Password = os.Getenv(constants.ElasticsearchPassword)
}
