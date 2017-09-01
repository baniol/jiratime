package main

import (
	"fmt"
	"os"

	"github.com/baniol/jiratime/commands"
	"github.com/mitchellh/cli"
)

var cmd map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	cmd = map[string]cli.CommandFactory{

		"days": func() (cli.Command, error) {
			return &commands.DaysCommand{
				Ui: ui,
			}, nil
		},

		"tickets": func() (cli.Command, error) {
			return &commands.TicketCommand{
				Ui: ui,
			}, nil
		},
	}
}

func main() {

	args := os.Args[1:]

	cli := &cli.CLI{
		Name:     "jiratime",
		Version:  "0.1.5",
		Args:     args,
		Commands: cmd,
	}

	_, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
	}

}
