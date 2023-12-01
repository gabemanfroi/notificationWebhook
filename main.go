package main

import (
	"github.com/gabemanfroi/notificationWebhook/domain/model"
	"github.com/gabemanfroi/notificationWebhook/internal/infra"
	ioc "github.com/gabemanfroi/notificationWebhook/internal/infra/IoC"
	"github.com/gabemanfroi/notificationWebhook/internal/infra/core"
	"github.com/gabemanfroi/notificationWebhook/internal/infra/data/queries"
	"github.com/gabemanfroi/notificationWebhook/internal/infra/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
)

type Payload struct {
	AlertId string `json:"alertId"`
}

func init() {
	log.Println("Setting up server...")
	if os.Getenv("APP_ENV") == "local" {
		utils.HandleError(godotenv.Load(), "Error loading .env file")
	}
	core.LoadConfig()
	ioc.InitContainer()
	log.Println("Server setup complete")
}

func main() {
	r := gin.Default()

	r.POST("/webhook/notify", func(c *gin.Context) {
		var payload Payload

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		queryJson := utils.GetQueryMarshalledJson(queries.GetAnalysisByAlertId(payload.AlertId))
		res := utils.ExecuteElasticsearchQuery(queryJson, core.GetElasticsearchClientInstance())

		var decodedResponse model.GetAnalysisByAlertIdResponse

		if err := utils.DecodeElasticsearchResponse(res, &decodedResponse); err != nil {
			if err == io.EOF {
				c.JSON(400, gin.H{"error": err.Error()})
			}
			c.JSON(400, gin.H{"error": err.Error()})
		}

		infra.NotifyAlertSlack(decodedResponse)
		infra.NotifyAlertTelegram(decodedResponse)

		c.JSON(200, decodedResponse)
	})

	err := r.Run(":8082")
	if err != nil {
		return
	}
}
