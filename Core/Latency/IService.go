package Latency

type IGetService interface {
	GetServiceAPITokenURL() string
	GetServiceAPIURL() string
	GetTargetURL() string
	GetRuns() int
	GetWaitInterval() int
	GetLocations() []string
	GetAPIKey() string
	GetOutputLocationsNumber() int
}

type ISetService interface {
	SetTargetURL(TargetUrl string)
	SetRuns(Runs int)
	SetWaitInterval(Interval int)
	SetLocations(Locations []string)
	SetOutputLocationsNumber(OutputLocationsNumber int)
	SetServiceAPITokenURL(Url string)
	SetServiceAPIURL(Url string)
}
