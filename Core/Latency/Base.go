package Latency

type LatencyChecker struct {
	TargetUrl             string
	Runs                  int
	WaitInterval          int
	Locations             []string
	APIKey                string
	ContentType           string
	OutputLocationsNumber int
	ServiceAPITokenURL    string
	ServiceAPIURL         string
}

type LatencyCheckerOutput struct {
	Location   string  `json:"location", yaml:"location"`
	AvgLatency float64 `json:"avgLatency", yaml:"avgLatency"`
}

type LatencyCheckerOutputList struct {
	Result []LatencyCheckerOutput `json:"result", yaml:"result"`
}

type TokenAPIResponse struct {
	RequestCount int `json:"request_count"`
	Duration     int `json:"duration"`
}

type LatencyAPIRequest struct {
	TargetURL string   `json:"target"`
	Locations []string `json:"locations"`
}
