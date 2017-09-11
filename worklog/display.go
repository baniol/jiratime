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

	for _, d := range days {
		v := perDay[d]
		hrs := v.Count / 3600
		logged += hrs
		v.Count = hrs
		filtered[d] = v
		fmt.Printf("%s\t%d\t%s\n", d, hrs, strings.Join(v.TicketKey[:], ", "))
	}

	required := len(filtered) * jiratimeConfig.HoursDaily

	fmt.Println("-----------------")
	fmt.Println("Logged\tRequired\t")
	fmt.Printf("%d\t%d\n", logged, required)
}
