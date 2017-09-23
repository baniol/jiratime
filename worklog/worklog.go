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

// JiraSession represents a struct with jiraClient
type JiraSession struct {
	Client *jira.Client
	Config *config.Config
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
func NewClient(config *config.Config) *JiraSession {
	jiraClient, err := jira.NewClient(nil, config.JiraURL)
	jiraClient.Authentication.SetBasicAuth(config.JiraUser, config.JiraPassword)
	if err != nil {
		panic(err)
	}
	return &JiraSession{Client: jiraClient, Config: config}
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
	jql := fmt.Sprintf("worklogDate>=%s AND worklogAuthor=%s", s.Config.DateFrom, s.Config.JiraUser)
	options := &jira.SearchOptions{MaxResults: s.Config.MaxSearchResults, Fields: []string{"worklog"}}
	tickets, _, err := s.Client.Issue.Search(jql, options)

	if err != nil {
		return nil, err
	}
	return tickets, nil
}

// MapHoursPerTicket aggregates logged hours by ticket
func MapHoursPerTicket(c *config.Config, tickets []jira.Issue) (HoursPerTicket, int) {
	var totalSpent int
	perTicket := make(HoursPerTicket, 0)
	for _, i := range tickets {
		worklogs := i.Fields.Worklog.Worklogs
		for _, w := range worklogs {
			// TODO: add tests checking author
			if w.Author.Name == c.JiraUser {
				perTicket[i.Key] += w.TimeSpentSeconds
				totalSpent += w.TimeSpentSeconds
			}
		}
	}
	return perTicket, totalSpent
}

// MapHoursPerDay aggregates logged hours by day
func MapHoursPerDay(c *config.Config, tickets []jira.Issue) HoursPerDay {
	perDay := make(HoursPerDay)
	for _, i := range tickets {
		worklogs := i.Fields.Worklog.Worklogs
		for _, w := range worklogs {
			// TODO: add tests checking author
			if w.Author.Name == c.JiraUser {
				t := time.Time(w.Started)
				f := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
				v := perDay[f]
				v.Count += w.TimeSpentSeconds
				ticket := ticketDetails{i.Key, w.TimeSpentSeconds}
				v.Ticket = append(v.Ticket, &ticket)
				perDay[f] = v
			}
		}
	}
	return perDay
}
