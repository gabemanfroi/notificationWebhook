package infra

import (
	"fmt"
	"github.com/gabemanfroi/notificationWebhook/domain/model"
	"github.com/slack-go/slack"
)

func NotifyAlert(alert model.GetAnalysisByAlertIdResponse) {
	SendSlackMessage(FormatSlackMessage(alert))
}

func SendSlackMessage(message string) {
	api := slack.New("xoxb-5463992198016-5825461165776-QPQjLvcaW8PFuK68puvomnU9")

	channelID, timestamp, err := api.PostMessage("C05PKJ9KPDH", slack.MsgOptionText(
		message, false), slack.MsgOptionAsUser(true))

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

func FormatSlackMessage(alert model.GetAnalysisByAlertIdResponse) string {
	return fmt.Sprintf("Você tem um novo alerta para o agente %s - %s: \n"+
		"Descrição %s \n"+
		"IP de origem: %s \n"+
		"Porta de origem: %s \n"+
		"IP de destino: %s \n"+
		"Porta de destino: %s \n"+
		"Protocolo: %s \n"+
		"Observável: %s \n",
		alert.Hits.Hits[0].Source.Alert.Agent.Name,
		alert.Hits.Hits[0].Source.Alert.Agent.Ip,
		alert.Hits.Hits[0].Source.Alert.Rule.Description,
		alert.Hits.Hits[0].Source.Alert.Data.SourceIp,
		alert.Hits.Hits[0].Source.Alert.Data.SourcePort,
		alert.Hits.Hits[0].Source.Alert.Data.DestinationIp,
		alert.Hits.Hits[0].Source.Alert.Data.DestinationPort,
		alert.Hits.Hits[0].Source.Alert.Data.Protocol,
		alert.Hits.Hits[0].Source.Analysis[0].Observable,
	)
}
