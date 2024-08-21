package Cli

import (
	"Taxocli/Config"
	"Taxocli/Core/Latency"
	"Taxocli/Core/Utils"
	"errors"

	"github.com/spf13/cobra"
)

func TocliAddExecFlags(cmd *cobra.Command) {
	Flags := cmd.Flags()
	Flags.StringVarP(&Config.TARGET_URL, "target-url", "t", "", "The target url. e.g: https://google.com")
	Flags.IntVarP(&Config.NUMBER_OF_RUNS, "runs", "r", 1, "The number of executions.")
	Flags.StringVarP(&Config.WAIT_INTERVAL, "interval", "i", "1m", "The amount of waiting time between runs.")
	Flags.StringArrayVarP(&Config.LOCATIONS, "locations", "l", []string{"EU.ES.*"}, "The array of locations to be requested. e.g: NA.US.*,NA.EU.*")
	Flags.IntVar(&Config.OUTPUT_LOCATIONS_NUMBER, "output-locations", 1, "The number of best locations to output.")
	Flags.StringVarP(&Config.OUTPUT_FORMAT, "output-format", "o", "table", "Output in an specific format. Usage: '-o [ table | yaml | json ]'")
	cmd.MarkFlagRequired("target-url")

}

func TocliNewExecCommand() *cobra.Command {

	Com := &cobra.Command{
		Use:   "run",
		Short: "Runtime The Run Command",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := TocliValidateCommandArgs()
			if err != nil {
				return err
			}
			WaitIntervalSeconds := Utils.IntervalTimeToSeconds(Config.WAIT_INTERVAL)
			this := Latency.NewLatencyChecker(Utils.GetEnv("TAXONIM_X_API_KEY", "NOT_SET"), Config.TARGET_URL, Config.NUMBER_OF_RUNS, WaitIntervalSeconds, Config.LOCATIONS, Config.OUTPUT_LOCATIONS_NUMBER)
			Response, err := this.RunCommandExec()

			switch {
			case Config.OUTPUT_FORMAT == "yaml":
				Utils.WriteOutputYaml(Response)
			case Config.OUTPUT_FORMAT == "json":
				Utils.WriteOutputJson(Response)
			default:
				Utils.WriteOutputTable(Response)

			}
			return err
		},
	}
	TocliAddExecFlags(Com)
	return Com
}

func TocliValidateCommandArgs() error {
	ValidInterval := Utils.ValidateIntervalTime(Config.WAIT_INTERVAL)
	if !ValidInterval {
		return errors.New(" [ ERROR ]: Invalid Interval Format")
	}
	ValidUrl := Utils.ValidateURL(Config.TARGET_URL)
	if !ValidUrl {
		return errors.New(" [ ERROR ]: Invalid Cron Format")
	}
	return nil

}
