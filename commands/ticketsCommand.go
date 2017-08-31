package commands

import (
	"fmt"
	"strings"

	"github.com/baniol/jiratime/config"
	"github.com/baniol/jiratime/worklog"
	"github.com/mitchellh/cli"
)

type TicketCommand struct {
	Ui         cli.Ui
	ShutdownCh <-chan struct{}
	args       []string
}

func (c *TicketCommand) Run(args []string) int {
	config, err := config.GetConfig()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error: %s", err))
		return 1
	}

	client := worklog.NewClient(&config)
	tickets := worklog.GetUserTickets(client)

	perTicket, total := worklog.MapHoursPerTicket(tickets)
	worklog.DisplayPerTicket(perTicket, total)
	return 0
}

func (c *TicketCommand) Synopsis() string {
	return "Displays a list of tickets with a number of hours logged"
}

func (c *TicketCommand) Help() string {
	helpText := `
Usage: jiratime tickets [options]

Displays a list of tickets with a number of hours logged

`
	return strings.TrimSpace(helpText)
}
