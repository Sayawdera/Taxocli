package Latency

import "Taxocli/Config"

func (this *LatencyChecker) GetServiceAPITokenURL() string {
	return this.APIKey
}

func (this *LatencyChecker) GetServiceAPIURL() string {
	return this.ServiceAPIURL
}
func (this *LatencyChecker) GetTargetURL() string {
	return this.TargetUrl
}
func (this *LatencyChecker) GetRuns() int {
	return this.Runs
}
func (this *LatencyChecker) GetWaitInterval() int {
	return this.WaitInterval
}
func (this *LatencyChecker) GetLocations() []string {
	return this.Locations
}
func (this *LatencyChecker) GetAPIKey() string {
	return this.APIKey
}
func (this *LatencyChecker) GetOutputLocationsNumber() int {
	return this.OutputLocationsNumber
}

func (this *LatencyChecker) SetTargetURL(TargetUrl string) {
	this.TargetUrl = TargetUrl
}
func (this *LatencyChecker) SetRuns(Runs int) {
	this.Runs = Runs
}
func (this *LatencyChecker) SetWaitInterval(Interval int) {
	this.WaitInterval = Interval
}
func (this *LatencyChecker) SetLocations(Locations []string) {
	this.Locations = Locations
}
func (this *LatencyChecker) SetOutputLocationsNumber(OutputLocationsNumber int) {
	this.OutputLocationsNumber = OutputLocationsNumber
}
func (this *LatencyChecker) SetServiceAPITokenURL(Url string) {
	this.ServiceAPITokenURL = Url
}
func (this *LatencyChecker) SetServiceAPIURL(Url string) {
	this.ServiceAPIURL = Url
}

func NewLatencyChecker(ApiKey string, TargetUrl string, Runs int, WaitInerval int, locations []string, OutputLocationsNumber int) *LatencyChecker {
	return &LatencyChecker{
		TargetUrl:             TargetUrl,
		Runs:                  Runs,
		WaitInterval:          WaitInerval,
		Locations:             locations,
		APIKey:                ApiKey,
		ContentType:           Config.TAXOCLI_CONTENT_TYPE_REQ,
		OutputLocationsNumber: OutputLocationsNumber,
		ServiceAPITokenURL:    Config.TAXOCLI_TOKEN_API_URL,
		ServiceAPIURL:         Config.TAXOCLI_LATENCY_API_URL,
	}
}
