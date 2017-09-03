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

var args = os.Args[1:]

func init() {

	// provide a default version string if app is built without makefile
	if VERSION == "" {
		VERSION = "version-manually-built"
	}

	ui := &cli.BasicUi{
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	meta := commands.Meta{
		Ui: ui,
	}

	cmd = map[string]cli.CommandFactory{

		"days": func() (cli.Command, error) {
			return &commands.DaysCommand{
				Meta: meta,
			}, nil
		},

		"tickets": func() (cli.Command, error) {
			return &commands.TicketCommand{
				Meta: meta,
			}, nil
		},
	}
}

func main() {

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
