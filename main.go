package main

import (
	"fmt"
	"os"

	"github.com/baniol/jiratime/commands"
	"github.com/mitchellh/cli"
)

// VERSION set with Makefile using linker flag, must be uninitialized
var VERSION string

var cmd map[string]cli.CommandFactory

func init() {

	// provide a default version string if app is built without makefile
	if VERSION == "" {
		VERSION = "version-manually-built"
	}

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
		Version:  VERSION,
		Args:     args,
		Commands: cmd,
	}

	_, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
	}

}
