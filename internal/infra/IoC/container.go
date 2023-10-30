package ioc

import (
	"github.com/gabemanfroi/notificationWebhook/internal/infra/core"
	"github.com/gabemanfroi/notificationWebhook/internal/infra/utils"
	"github.com/golobby/container/v3"
	"github.com/opensearch-project/opensearch-go"
)

func InitContainer() {
	bindSingletons()
}

func bindSingletons() {
	utils.HandleError(container.Singleton(func() *opensearch.Client { return core.CreateElasticsearchClient() }),
		"Error while binding singletons - [ElasticsearchClient]")
}
