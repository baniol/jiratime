package commands

import (
	"strings"

	"github.com/baniol/jiratime/worklog"
)

type TicketCommand struct {
	Meta
}

func (c *TicketCommand) Run(args []string) int {

	cmdFlags := c.FlagSet("days")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	configPath := cmdFlags.Lookup("config").Value.String()
	config, err := c.Config(configPath)
	if err != nil {
		return 1
	}

	client := worklog.NewClient(config)
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
