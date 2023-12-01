package infra

import (
	"fmt"
	"github.com/gabemanfroi/notificationWebhook/domain/model"
	"github.com/gabemanfroi/notificationWebhook/internal/infra/core"
	"github.com/slack-go/slack"
)

func NotifyAlertSlack(alert model.GetAnalysisByAlertIdResponse) {
	SendSlackMessage(FormatMessage(alert))
}

func SendSlackMessage(message string) {
	api := slack.New(core.AppConfig.Slack.Token)

	channelID, timestamp, err := api.PostMessage("C05PKJ9KPDH", slack.MsgOptionText(
		message, false), slack.MsgOptionAsUser(true))

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
