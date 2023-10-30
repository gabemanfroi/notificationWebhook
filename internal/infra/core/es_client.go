package core

import (
	"crypto/tls"
	"github.com/gabemanfroi/notificationWebhook/internal/infra/utils"
	"github.com/golobby/container/v3"
	"github.com/opensearch-project/opensearch-go"
	"io/ioutil"
	"net/http"
)

func CreateElasticsearchClient() *opensearch.Client {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	caCert, err := ioutil.ReadFile("./auth.pem")

	config := opensearch.Config{
		Addresses: []string{AppConfig.Elasticsearch.Host},
		Username:  AppConfig.Elasticsearch.Username,
		Password:  AppConfig.Elasticsearch.Password,
		Transport: transport,
		CACert:    caCert,
	}

	client, err := opensearch.NewClient(config)
	if err != nil {
		panic(err)
	}

	return client

}

func GetElasticsearchClientInstance() *opensearch.Client {
	var injected *opensearch.Client
	utils.HandleError(container.Resolve(&injected), "Error while resolving elasticsearch client")
	return injected
}
