package core

import (
	"log"
	"os"
)

type Config struct {
	Elasticsearch struct {
		Host     string
		Username string
		Password string
	}
	Telegram struct {
		Token string
	}
	Slack struct {
		Token string
	}
	Wazuh struct {
		AlertsIndex string
	}
	App struct {
		Environment string
	}
}

var AppConfig = Config{}

func LoadConfig() {
	log.Println("Loading config...")
	AppConfig.App.Environment = os.Getenv(AppEnvironment)
	AppConfig.Elasticsearch.Host = os.Getenv(ElasticsearchHost)
	AppConfig.Elasticsearch.Username = os.Getenv(ElasticsearchUsername)
	AppConfig.Elasticsearch.Password = os.Getenv(ElasticsearchPassword)
	AppConfig.Telegram.Token = os.Getenv(TelegramToken)
	AppConfig.Wazuh.AlertsIndex = os.Getenv(WazuhAlertsIndex)
	AppConfig.Slack.Token = os.Getenv(SlackToken)
}
