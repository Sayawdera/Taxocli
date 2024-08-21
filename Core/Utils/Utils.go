package Utils

import (
	"Taxocli/Core/Latency"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"

	color "github.com/TwiN/go-color"
	"github.com/adhocore/gronx"
	"github.com/jedib0t/go-pretty/v6/table"
)

func GetEnv(Input string, Output string) string {
	if VarValue, ok := os.LookupEnv(Input); ok {
		return VarValue
	}
	return Output
}

func ValidateURL(InputUrl string) bool {
	Url, err := url.Parse(InputUrl)
	return err == nil && Url.Scheme != "" && Url.Host != ""
}

func ValidateIntervalTime(Interval string) bool {
	Reg := regexp.MustCompile(`^(\d*)(s|m|h)`)
	Matched := Reg.MatchString(Interval)
	return Matched
}

func ValidateCronTime(CronTime string) bool {
	log.Println(" [ INFO ]: CronTime To Be Validated", CronTime)
	Gron := gronx.New()
	if CronTime != "" {
		return Gron.IsValid(CronTime)
	}
	return false
}

func GetNextTimeCronTime(CronTime string) int64 {
	NextTime, err := gronx.NextTick(CronTime, false)
	if err != nil {
		log.Println(" [ INFO ]: CronTime Format Is Not Valid Or Not Scheduled CronTime Passed", err)
		return -1
	}
	Duration := NextTime.Unix() - time.Now().Unix()
	return Duration
}

func IntervalTimeToSeconds(Interval string) int {
	Reg := regexp.MustCompile(`^(\d*)(s|m|h)`)
	CaptureGroups := Reg.FindStringSubmatch(Interval)
	if len(CaptureGroups) < 1 {
		return -1
	}
	TimeValue, err := strconv.Atoi(CaptureGroups[1])
	if err != nil {
		return -1
	}
	TimeUnit := CaptureGroups[2]

	switch TimeUnit {
	case "s":
		return TimeValue
	case "m":
		return TimeValue * 60
	case "h":
		return TimeValue * 60 * 60
	}
	return -1
}

func WriteOutputTable(this Latency.LatencyCheckerOutputList) {
	T := table.NewWriter()
	T.SetOutputMirror(os.Stdout)
	T.AppendHeader(table.Row{color.InWhite("Location"), color.InWhite("Average Latency")})

	if len(this.Result) > 0 {
		for A := range this.Result {
			if A == 0 {
				T.AppendRow([]interface{}{color.InGreen(this.Result[A].Location), color.InGreen(fmt.Sprintf("{ %f }", this.Result[A].AvgLatency))})
				continue
			}
			T.AppendRow([]interface{}{color.InWhite(this.Result[A].Location), color.InWhite(fmt.Sprintf("%f", this.Result[A].AvgLatency))})
		}
	}
	T.SetStyle(table.StyleLight)
	T.Render()
}

func WriteOutputJson(this Latency.LatencyCheckerOutputList) {
	Output, _ := json.MarshalIndent(this, "", "  ")
	fmt.Println(string(Output))
}
func WriteOutputYaml(this Latency.LatencyCheckerOutputList) {
	Output, _ := yaml.Marshal(this)
	fmt.Println(string(Output))
}
