package tests

import (
	"Taxocli/Core/Latency"
	"Taxocli/Core/Utils"
	"reflect"
	"testing"
)

func TestIntervalTImeToSeconds(t *testing.T) {
	type Args struct {
		Interval string
	}
	tests := []struct {
		Name string
		Args Args
		Want int
	}{
		{
			Name: "Test OK Testing minutes",
			Args: Args{
				Interval: "1m",
			},
			Want: 60,
		},

		{
			Name: "Test KO testing Wrong Conversion Type",
			Args: Args{
				Interval: "TestKO",
			},
			Want: -1,
		},
		{
			Name: "Test KO wrong duration type",
			Args: Args{
				Interval: "1q",
			},
			Want: -1,
		},
		{
			Name: "Test KO testing seconds",
			Args: Args{
				Interval: "4s",
			},
			Want: 4,
		},
		{
			Name: "Test KO testing hours",
			Args: Args{
				Interval: "1h",
			},
			Want: 3600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if Got := Utils.IntervalTimeToSeconds(tt.Args.Interval); Got != tt.Want {
				t.Errorf("IntervalTImeToSeconds() = %v, want %v", Got, tt.Want)
			}
		})
	}
}

func TestLatencyChecherGetMinimumLatencies(t *testing.T) {

	type Fields struct {
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
	type Args struct {
		Latencies map[string]float64
	}
	tests := []struct {
		Name      string
		Fields    Fields
		Args      Args
		Want      []string
		WantError []float64
	}{}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			this := &Latency.LatencyChecker{
				TargetUrl:             tt.Fields.TargetUrl,
				Runs:                  tt.Fields.Runs,
				WaitInterval:          tt.Fields.WaitInterval,
				Locations:             tt.Fields.Locations,
				APIKey:                tt.Fields.APIKey,
				ContentType:           tt.Fields.ContentType,
				OutputLocationsNumber: tt.Fields.OutputLocationsNumber,
				ServiceAPITokenURL:    tt.Fields.ServiceAPITokenURL,
				ServiceAPIURL:         tt.Fields.ServiceAPIURL,
			}
			Ram, Rem := this.GetMinimumLatencies(tt.Args.Latencies)
			if !reflect.DeepEqual(Ram, tt.Want) {
				t.Errorf("getMinimumLatencies() got = %v, want %v", Ram, tt.Want)
			}
			if !reflect.DeepEqual(Rem, tt.Want) {
				t.Errorf("getMinimumLatencies() got1 = %v, want %v", Rem, tt.Want)
			}
		})
	}
}

func TestValidateIntervalTime(t *testing.T) {
	type Args struct {
		Interval string
	}
	tests := []struct {
		Name string
		Args Args
		Want bool
	}{
		{
			Name: "Test OK Testing Validation Time Seconds",
			Args: Args{
				Interval: "1s",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Time Minutes",
			Args: Args{
				Interval: "1m",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Time Hours",
			Args: Args{
				Interval: "1h",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Time Type",
			Args: Args{
				Interval: "1q",
			},
			Want: false,
		},
		{
			Name: "Test OK Testing Validation Time Empty",
			Args: Args{
				Interval: "",
			},
			Want: false,
		},
		{
			Name: "Test OK Testing Validation Time Wron Value",
			Args: Args{
				Interval: "kk",
			},
			Want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if Got := Utils.ValidateIntervalTime(tt.Args.Interval); Got != tt.Want {
				t.Errorf("ValidateIntervalTime() = %v, want %v", Got, tt.Want)
			}
		})
	}
}

func TestValidateUrl(t *testing.T) {
	type Args struct {
		Url string
	}
	tests := []struct {
		Name string
		Args Args
		Want bool
	}{
		{
			Name: "Test OK Testing Validation Url",
			Args: Args{
				Url: "https://test.com",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Url",
			Args: Args{
				Url: "http://test.com",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Url",
			Args: Args{
				Url: "http://test.com",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Url",
			Args: Args{
				Url: "https://test.com",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Url",
			Args: Args{
				Url: "http://test.com",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Url",
			Args: Args{
				Url: "https://test.com",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Url",
			Args: Args{
				Url: "http://test.com",
			},
			Want: true,
		},
		{
			Name: "Test OK Testing Validation Url",
			Args: Args{
				Url: "https://test.com",
			},
			Want: true,
		},
	}
	for _, TT := range tests {
		t.Run(TT.Name, func(t *testing.T) {
			if Got := Utils.ValidateURL(TT.Args.Url); Got != TT.Want {
				t.Errorf("ValidateUrl() = %v, want %v", Got, TT.Want)
			}
		})
	}
}
