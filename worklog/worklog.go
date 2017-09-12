package worklog

import (
	"fmt"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/baniol/jiratime/config"
)

type jiraGetter interface {
	getUserWorklog() ([]jira.Issue, error)
}

var jiratimeConfig *config.Config

// JiraSession represents a struct with jiraClient
type JiraSession struct {
	Client *jira.Client
}

// HoursPerDay maps logged hours to days
type HoursPerDay map[string]struct {
	Count  int
	Ticket []*ticketDetails
}

type ticketDetails struct {
	Key   string
	Hours int
}

// HoursPerTicket maps logged hours per ticket
type HoursPerTicket map[string]int

// NewClient returns an instance on JiraSession containing jiraClient.
// It sets global jiratimeConfig with passed config parameter.
func NewClient(config *config.Config) *JiraSession {
	jiraClient, err := jira.NewClient(nil, config.JiraURL)
	jiraClient.Authentication.SetBasicAuth(config.JiraUser, config.JiraPassword)
	if err != nil {
		panic(err)
	}
	jiratimeConfig = config
	return &JiraSession{Client: jiraClient}
}

// GetUserTickets is an exported function and returns a list of user worklogs
func GetUserTickets(c jiraGetter) []jira.Issue {
	items, err := c.getUserWorklog()
	if err != nil {
		panic(err)
	}
	return items
}

// Worklog returns a list of Issues where a particular user saved working hours
func (s *JiraSession) getUserWorklog() ([]jira.Issue, error) {
	jql := fmt.Sprintf("worklogDate>=%s AND worklogAuthor=%s", jiratimeConfig.DateFrom, jiratimeConfig.JiraUser)
	options := &jira.SearchOptions{MaxResults: jiratimeConfig.MaxSearchResults, Fields: []string{"worklog"}}
	tickets, _, err := s.Client.Issue.Search(jql, options)

	if err != nil {
		return nil, err
	}
	return tickets, nil
}

// MapHoursPerTicket aggregates logged hours by ticket
func MapHoursPerTicket(tickets []jira.Issue) (HoursPerTicket, int) {
	var totalSpent int
	perTicket := make(HoursPerTicket, 0)
	for _, i := range tickets {
		worklogs := i.Fields.Worklog.Worklogs
		for _, w := range worklogs {
			perTicket[i.Key] += w.TimeSpentSeconds
			totalSpent += w.TimeSpentSeconds
		}
	}
	return perTicket, totalSpent
}

// MapHoursPerDay aggregates logged hours by day
func MapHoursPerDay(tickets []jira.Issue) HoursPerDay {
	perDay := make(HoursPerDay)
	for _, i := range tickets {
		worklogs := i.Fields.Worklog.Worklogs
		for _, w := range worklogs {
			t := time.Time(w.Started)
			f := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
			v := perDay[f]
			v.Count += w.TimeSpentSeconds
			ticket := ticketDetails{i.Key, w.TimeSpentSeconds}
			v.Ticket = append(v.Ticket, &ticket)
			perDay[f] = v
		}
	}
	return perDay
}
