package tests

import (
	"Taxocli/Config"
	"Taxocli/Core/Latency"
	"reflect"
	"testing"
)

func TestLatencyCheckerGetAPIKey(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test OK get case Get APIKEY",
			fields: fields{
				TargetUrl:             "",
				Runs:                  0,
				WaitInterval:          0,
				Locations:             nil,
				APIKey:                "TESTOK",
				ContentType:           "",
				OutputLocationsNumber: 0,
				ServiceAPITokenURL:    "",
				ServiceAPIURL:         "",
			},
			want: "TESTOK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			if got := lc.GetAPIKey(); got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatencyCheckerGetLocations(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			if got := lc.GetLocations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatencyCheckerGetOutputLocationsNumber(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			if got := lc.GetOutputLocationsNumber(); got != tt.want {
				t.Errorf("GetOutputLocationsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatencyCheckerGetRuns(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			if got := lc.GetRuns(); got != tt.want {
				t.Errorf("GetRuns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatencyCheckerGetServiceAPITokenURL(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			if got := lc.GetServiceAPITokenURL(); got != tt.want {
				t.Errorf("GetServiceAPITokenURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatencyCheckerGetServiceAPIURL(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			if got := lc.GetServiceAPIURL(); got != tt.want {
				t.Errorf("GetServiceAPIURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatencyCheckerGetTargetURL(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			if got := lc.GetTargetURL(); got != tt.want {
				t.Errorf("GetTargetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatencyCheckerGetWaitInterval(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			if got := lc.GetWaitInterval(); got != tt.want {
				t.Errorf("GetWaitInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLatencyCheckerSetLocations(t *testing.T) {
	type fields struct {
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
	type args struct {
		locations []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TEST OK SetLocations",
			fields: fields{
				TargetUrl:             "",
				Runs:                  0,
				WaitInterval:          0,
				Locations:             []string{"us-east-1", "us-west-1"},
				APIKey:                "",
				ContentType:           "",
				OutputLocationsNumber: 0,
				ServiceAPITokenURL:    "",
				ServiceAPIURL:         "",
			},
			args: args{
				locations: []string{"us-east-1", "us-west-1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			lc.SetLocations(tt.args.locations)
		})
	}
}

func TestLatencyCheckerSetOutputLocationsNumber(t *testing.T) {
	type fields struct {
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
	type args struct {
		outputLocationsNumber int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			lc.SetOutputLocationsNumber(tt.args.outputLocationsNumber)
		})
	}
}

func TestLatencyCheckerSetRuns(t *testing.T) {
	type fields struct {
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
	type args struct {
		runs int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			lc.SetRuns(tt.args.runs)
		})
	}
}

func TestLatencyCheckerSetServiceAPITokenURL(t *testing.T) {
	type fields struct {
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
	type args struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			lc.SetServiceAPITokenURL(tt.args.URL)
		})
	}
}

func TestLatencyCheckerSetServiceAPIURL(t *testing.T) {
	type fields struct {
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
	type args struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			lc.SetServiceAPIURL(tt.args.URL)
		})
	}
}

func TestLatencyCheckerSetTargetURL(t *testing.T) {
	type fields struct {
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
	type args struct {
		targetUrl string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			lc.SetTargetURL(tt.args.targetUrl)
		})
	}
}

func TestLatencyCheckerSetWaitInterval(t *testing.T) {
	type fields struct {
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
	type args struct {
		interval int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Latency.LatencyChecker{
				TargetUrl:             tt.fields.TargetUrl,
				Runs:                  tt.fields.Runs,
				WaitInterval:          tt.fields.WaitInterval,
				Locations:             tt.fields.Locations,
				APIKey:                tt.fields.APIKey,
				ContentType:           tt.fields.ContentType,
				OutputLocationsNumber: tt.fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.fields.ServiceAPIURL,
			}
			lc.SetWaitInterval(tt.args.interval)
		})
	}
}

func TestNewLatencyChecker(t *testing.T) {
	type args struct {
		targetURL             string
		runs                  int
		waitInterval          int
		locations             []string
		outputLocationsNumber int
	}
	tests := []struct {
		name string
		args args
		want *Latency.LatencyChecker
	}{
		{
			name: "Test OK Cosntructor",
			args: args{
				targetURL:             "Test",
				runs:                  1,
				waitInterval:          1,
				locations:             nil,
				outputLocationsNumber: 1,
			},
			want: &Latency.LatencyChecker{
				TargetUrl:             "Test",
				Runs:                  1,
				WaitInterval:          1,
				Locations:             nil,
				APIKey:                "NOT_SET",
				ContentType:           Config.TAXOCLI_CONTENT_TYPE_REQ,
				OutputLocationsNumber: 1,
				ServiceAPITokenURL:    Config.TAXOCLI_TOKEN_API_URL,
				ServiceAPIURL:         Config.TAXOCLI_LATENCY_API_URL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Latency.NewLatencyChecker("NOT_SET", tt.args.targetURL, tt.args.runs, tt.args.waitInterval, tt.args.locations, tt.args.outputLocationsNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLatencyChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
