package calendar

import (
	"fmt"
	"time"
)

// Days represents a list of working days in a month
type Days []string

// GetDays return a list of working days of a given month
func GetDays(year int, month int) Days {

	days := make(Days, 0)

	if year == 0 || month == 0 {
		currentTime := time.Now().Local()
		if year == 0 {
			year = currentTime.Year()
		}
		if month == 0 {
			month = int(currentTime.Month())
		}
	}

	date := fmt.Sprintf("2017-%d-1", month)
	start, _ := time.Parse("2006-1-2", date)

	for t := start; t.Month() == start.Month(); t = t.AddDate(0, 0, 1) {
		name := t.Weekday().String()
		if name != "Saturday" && name != "Sunday" {
			full := fmt.Sprintf("%d-%02d-%02d", year, int(t.Month()), t.Day())
			days = append(days, full)
		}
	}

	return days
}

// PrepareDateParams checks if the inpur year and month have proper values.
// If the proper value is not passed, it returns current year and month.
func PrepareDateParams(yearParam int, monthParam int) (int, int) {
	year := yearParam
	month := monthParam

	t := time.Now().Local()

	if yearParam == 0 {
		year = t.Year()
	}

	if monthParam == 0 {
		month = int(t.Month())
	}

	return year, month
}
