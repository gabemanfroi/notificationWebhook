package infra

import (
	"fmt"
	"github.com/gabemanfroi/notificationWebhook/domain/model"
)

func FormatMessage(alert model.GetAnalysisByAlertIdResponse) string {
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
