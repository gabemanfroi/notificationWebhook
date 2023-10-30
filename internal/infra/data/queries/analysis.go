package queries

func GetAnalysisByAlertId(alertId string) map[string]interface{} {
	return map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"alert.id": alertId,
						},
					},
				},
			},
		},
		"_source": map[string]interface{}{
			"includes": []string{
				"alert", "analysis",
			},
		},
	}
}
