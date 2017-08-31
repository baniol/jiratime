package commands

import (
	"flag"
	"fmt"
	"strings"

	"github.com/baniol/jiratime/calendar"
	"github.com/baniol/jiratime/config"
	"github.com/baniol/jiratime/worklog"
	"github.com/mitchellh/cli"
)

type DaysCommand struct {
	Ui cli.Ui
}

func (c *DaysCommand) Run(args []string) int {
	config, err := config.GetConfig()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error: %s", err))
		return 1
	}

	cmdFlags := flag.NewFlagSet("fmt", flag.ContinueOnError)
	var yearParam int
	var monthParam int
	cmdFlags.IntVar(&yearParam, "year", 0, "Year to display")
	cmdFlags.IntVar(&monthParam, "month", 0, "Month to display")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	client := worklog.NewClient(&config)
	tickets := worklog.GetUserTickets(client)

	year, month := calendar.PrepareDateParams(yearParam, monthParam)

	hoursPerDay := worklog.MapHoursPerDay(tickets)
	worklog.DisplayPerMonth(hoursPerDay, year, month)

	return 0
}

func (c *DaysCommand) Synopsis() string {
	return "Displays a list of days of a month with a number of hours logged"
}

func (c *DaysCommand) Help() string {
	helpText := `
Usage: jiratime days [options]

Displays a list of days of a month with a number of hours logged

Options:

  -year  <int>		Specify a year to be displayed. Default - current year.
  -month <int>		Specify a month to be displayed. Default - current month.

`
	return strings.TrimSpace(helpText)
}
