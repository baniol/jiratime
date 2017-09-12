package worklog

import (
	"fmt"
	"github.com/baniol/jiratime/calendar"
	"strings"
)

// DisplayPerTicket outputs the result to stdout
func DisplayPerTicket(tickets HoursPerTicket, total int) {

	fmt.Print("Ticket\tHours\n")
	for k, e := range tickets {
		fmt.Printf("%s\t%d\n", k, e/3600)
	}
	fmt.Println("----------------------")
	fmt.Printf("Total logged: %d hours\n", total/3600)
}

// DisplayPerMonth outputs the result to stdout
func DisplayPerMonth(perDay HoursPerDay, year int, month int) {

	fmt.Print("Day\t\tHours\tTickets\n")

	logged := 0
	days := calendar.GetDays(year, month)

	filtered := make(HoursPerDay)

	// @TODO refactor
	for _, d := range days {
		v := perDay[d]
		hrs := v.Count / 3600
		logged += hrs
		v.Count = hrs
		filtered[d] = v

		ticketDetails := make([]string, 0)
		for _, t := range v.Ticket {
			ticketDetails = append(ticketDetails, fmt.Sprintf("%s (%d)", t.Key, t.Hours/3600))
		}

		fmt.Printf("%s\t%d\t%s\n", d, hrs, strings.Join(ticketDetails[:], ","))
	}

	required := len(filtered) * jiratimeConfig.HoursDaily

	fmt.Println("-----------------")
	fmt.Println("Logged\tRequired\t")
	fmt.Printf("%d\t%d\n", logged, required)
}
