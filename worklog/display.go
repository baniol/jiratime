package worklog

import (
	"fmt"
	"github.com/baniol/jiratime/calendar"
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

	fmt.Print("Day\t\tHours\n")

	logged := 0
	days := calendar.GetDays(year, month)

	filtered := make(HoursPerDay, 0)

	for _, d := range days {
		hrs := perDay[d] / 3600
		logged += hrs
		filtered[d] = hrs
		fmt.Printf("%s\t%d\n", d, hrs)
	}

	required := len(filtered) * jiraConfig.HoursDaily

	fmt.Println("-----------------")
	fmt.Println("Logged\tRequired\t")
	fmt.Printf("%d\t%d\n", logged, required)
}
