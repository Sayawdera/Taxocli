package main

import (
	"Taxocli/Core/Cli"
	color "github.com/TwiN/go-color"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	command := TocliNewCommands()
	if err := command.Execute(); err != nil {
		log.Fatalf(color.InRed("[ ERROR ]: ")+"%s", err.Error())
	}
}

func TocliNewCommands() *cobra.Command {
	this := &cobra.Command{
		Use:   "Taxocli-Latencies",
		Short: "Taxocli Latencies is The Command Line Interface To Work With The Taxonim Latencies API",
		Run: func(cmd *cobra.Command, Args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	this.AddCommand(Cli.TocliNewExecCommand())

	return this
}
