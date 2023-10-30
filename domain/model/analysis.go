package model

type GetAnalysisByAlertIdResponse struct {
	Hits struct {
		Hits []struct {
			Source struct {
				Alert struct {
					Id    string `json:"id"`
					Agent struct {
						Ip   string `json:"ip"`
						Name string `json:"name"`
					} `json:"agent"`
					Data struct {
						SourceIp        string `json:"src_ip"`
						DestinationIp   string `json:"dest_ip"`
						SourcePort      string `json:"src_port"`
						DestinationPort string `json:"dest_port"`
						Protocol        string `json:"proto"`
					} `json:"data"`
					Rule struct {
						Description string `json:"description"`
					} `json:"rule"`
				} `json:"alert"`
				Analysis []struct {
					Observable string `json:"observable"`
					Report     struct {
						Summary struct {
							Level     string `json:"level"`
							Namespace string `json:"namespace"`
							Predicate string `json:"predicate"`
							Value     string `json:"value"`
						} `json:"summary"`
					} `json:"report"`
				} `json:"analysis"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
