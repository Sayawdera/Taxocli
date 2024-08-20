package tests

//goland:noinspection ALL
import (
	"Taxocli/Config"
	"Taxocli/Core/Latency"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLatencyCheckerRUnCommandExec(t *testing.T) {
	var (
		Input        Latency.LatencyAPIRequest
		OutputTokeOK = Latency.TokenAPIResponse{
			RequestCount: 100000,
			Duration:     100,
		}
		OutputTokenKOInsufficient = Latency.TokenAPIResponse{
			RequestCount: 1,
			Duration:     100,
		}
		OutputLatencyOK = Latency.LatencyCheckerOutputList{
			Result: []Latency.LatencyCheckerOutput{
				Location:   "us-east-1",
				AvgLatency: 200,
			},
		}
	)

	Loc := []string{"us-east-1"}
	LocLocationDown := []string{"LocationDown"}
	Srv := httptest.NewServer(http.HandlerFunc(func(W http.ResponseWriter, R *http.Request) {
		if R.URL.Path == "/token" {
			W.WriteHeader(200)
			OutputTok, _ := json.Marshal(OutputTokeOK)
			W.Write(OutputTok)
			return
		}

		if R.URL.Path == "/tokenInssuficient" {
			W.WriteHeader(200)
			Output, _ := json.Marshal(OutputTokenKOInsufficient)
			W.Write(Output)
			return
		}

		if R.URL.Path == "/tokenBadStatusCode" {
			W.WriteHeader(400)
			Output, _ := json.Marshal(OutputTokenKOInsufficient)
			W.Write(Output)
			return
		}

		if R.URL.Path == "/tokenBadJsonResponse" {
			W.WriteHeader(400)
			Output, _ := json.Marshal("{OutputTokenKOInsufficient}")
			W.Write(Output)
			return
		}
		json.NewDecoder(R.Body).Decode(&Input)

		if Input.TargetURL == "https://requestOK.test" {

			if Input.Locations[0] == Loc[0] {
				W.WriteHeader(200)
				JsonResponse := `{"us-east-1":{"avgLatency":200, "StatusCode": 200}}`
				W.Write([]byte(JsonResponse))
				return
			}

			if Input.Locations[0] == LocLocationDown[0] {
				W.WriteHeader(200)
				JsonResponse := `{"us-east-1":{"avgLatency":200, "StatusCode": 400}}`
				W.Write([]byte(JsonResponse))
				return
			}
			W.Write([]byte(" [ ERROR ]: "))
			W.WriteHeader(400)
			return
		}
	}))
	defer Srv.Close()

	type Fields struct {
		TargetURL string
		Runs int
		WaitInterval int
		Locations []string
		APIKey string
		ContentType string
		OutputLocationsNumber int
		ServiceAPITokenURL string
		ServiceAPIURL string
	}
	tests := []struct {
		Name string
		Fields Fields
		Want Latency.LatencyCheckerOutputList
		WantError bool
	} {
		{
			Name: "Test OK Executing The Function With No Errors",
			Fields: Fields{
				TargetURL: "https://requestOK.test",
				Runs: 1,
				WaitInterval: 1,
				Locations: []string{"us-east-1", "us-west-1"},
				APIKey: "API KEY",
				ContentType: Config.TAXOCLI_CONTENT_TYPE_REQ,
				OutputLocationsNumber: 1,
				ServiceAPITokenURL: Srv.URL + "/token",
				ServiceAPIURL: Srv.URL,
			},
			Want: OutputLatencyOK,
			WantError: false,
		},
		{
			Name: "Test OK Executing The Function With No Errors, Multiple Runs",
			Fields: Fields{
				TargetURL: "https://requestOK.test",
				Runs: 2,
				WaitInterval: 1,
				Locations: []string{"us-east-1", "us-west-1"},
				APIKey: "API KEY",
				ContentType: Config.TAXOCLI_CONTENT_TYPE_REQ,
				OutputLocationsNumber: 1,
				ServiceAPITokenURL: Srv.URL + "/token",
				ServiceAPIURL: Srv.URL,
			},
			Want: OutputLatencyOK,
			WantError: false,
		},
		{
			Name: "Test KO To Test Errors In AvailableTokens",
			Fields: Fields{
				TargetURL: "https://requestOK.test",
				Runs: 1,
				WaitInterval: 1,
				Locations: []string{"us-east-1", "us-west-1"},
				APIKey: "API KEY",
				ContentType: Config.TAXOCLI_CONTENT_TYPE_REQ,
				OutputLocationsNumber: 1,
				ServiceAPITokenURL: Srv.URL + "/token",
				ServiceAPIURL: Srv.URL,
			},
			Want: Latency.LatencyCheckerOutputList{},
			WantError: false,
		},
		{
			Name: "Test KO To Test Errors In AvailableTokens",
			Fields: Fields{
				TargetURL: "https://requestKO.test",
				Runs: 1,
				WaitInterval: 1,
				Locations: []string{"us-east-1", "us-west-1"},
				APIKey: "API KEY",
				ContentType: Config.TAXOCLI_CONTENT_TYPE_REQ,
				OutputLocationsNumber: 1,
				ServiceAPITokenURL: Srv.URL + "/token",
				ServiceAPIURL: Srv.URL,
			},
			Want: Latency.LatencyCheckerOutputList{},
			WantError: true,
		},
		{
			Name: "Test OK Executing The Function With No Errors",
			Fields: Fields{
				TargetURL: "https://requestOK.test",
				Runs: 1,
				WaitInterval: 1,
				Locations: []string{"us-east-1", "us-west-1"},
				APIKey: "API KEY",
				ContentType: Config.TAXOCLI_CONTENT_TYPE_REQ,
				OutputLocationsNumber: 1,
				ServiceAPITokenURL: Srv.URL + "/token",
				ServiceAPIURL: Srv.URL,
			},
			Want: Latency.LatencyCheckerOutputList{},
			WantError: false,
		},
		{
			Name: "Test OK Executing The Function With No Errors",
			Fields: Fields{
				TargetURL: "https://requestOK.test",
				Runs: 1,
				WaitInterval: 1,
				Locations: []string{"us-east-1", "us-west-1"},
				APIKey: "API KEY",
				ContentType: Config.TAXOCLI_CONTENT_TYPE_REQ,
				OutputLocationsNumber: 1,
				ServiceAPITokenURL: Srv.URL + "/token",
				ServiceAPIURL: Srv.URL,
			},
			Want: Latency.LatencyCheckerOutputList{},
			WantError: false,
		},
		{
			Name: "Test OK Executing The Function With No Errors",
			Fields: Fields{
				TargetURL: "https://requestOK.test",
				Runs: 1,
				WaitInterval: 1,
				Locations: []string{"us-east-1", "us-west-1"},
				APIKey: "API KEY",
				ContentType: Config.TAXOCLI_CONTENT_TYPE_REQ,
				OutputLocationsNumber: 1,
				ServiceAPITokenURL: Srv.URL + "/token",
				ServiceAPIURL: Srv.URL,
			},
			Want: Latency.LatencyCheckerOutputList{},
			WantError: false,
		}

		for _, TT := range tests {
			t.Run(TT.Name, func(t *testing.T) {
				this := &LatencyChecker{
					TargetUrl:             TT.FIelds.TargetURL,
					Runs:                  TT.FIelds.Runs,
					WaitInterval:          TT.FIelds.WaitInterval,
					Locations:             TT.FIelds.Locations,
					APIKey:                TT.FIelds.APIKey,
					ContentType:           TT.FIelds.ContentType,
					OutputLocationsNumber: TT.FIelds.OutputLocationsNumber,
					ServiceAPITokenURL:    TT.FIelds.ServiceAPITokenURL,
					ServiceAPIURL:         TT.FIelds.ServiceAPIURL,
				}
				log.Println(this.locations)
				Got, err := this.RunCommandExec(
					if (err != nil) != TT.WantError {
						t.Errorf("RunCommandExec() Error = { %v } WatnError { %v }", err, TT.WantError)
					}
					if !reflect.DeepEqual(Got, TT.Want) {
						t.Errrof("RunCommandExec() = { %v } Want { %v }", Got, TT.Want)
					}
				)
			})
		}
	},
}
