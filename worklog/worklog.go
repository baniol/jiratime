package worklog

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/baniol/jiratime/config"
	"time"
)

type jiraGetter interface {
	getUserWorklog() ([]jira.Issue, error)
}

type jiraSession struct {
	Client *jira.Client
}

// HoursPerDay maps logged hours to days
type HoursPerDay map[string]int

// HoursPerTicket maps logged hours per ticket
type HoursPerTicket map[string]int

var jiraConfig *config.Config

func NewClient(config *config.Config) *jiraSession {
	jiraClient, err := jira.NewClient(nil, config.JiraURL)
	jiraClient.Authentication.SetBasicAuth(config.JiraUser, config.JiraPassword)
	if err != nil {
		panic(err)
	}
	// @TODO does setting the config belong to the function ?
	jiraConfig = config
	return &jiraSession{Client: jiraClient}
}

// GetUserTickets is an expored function and returns a list of user worklogs
func GetUserTickets(c jiraGetter) []jira.Issue {
	items, err := c.getUserWorklog()
	if err != nil {
		panic(err)
	}
	return items
}

// Worklog returns a list of Issues where a particular user saved working hours
func (s *jiraSession) getUserWorklog() ([]jira.Issue, error) {
	jql := fmt.Sprintf("worklogDate>=%s AND worklogAuthor=%s", jiraConfig.DateFrom, jiraConfig.JiraUser)
	options := &jira.SearchOptions{MaxResults: jiraConfig.MaxSearchResults, Fields: []string{"worklog"}}
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
	perDay := make(HoursPerDay, 0)
	for _, i := range tickets {
		worklogs := i.Fields.Worklog.Worklogs
		for _, w := range worklogs {
			t := time.Time(w.Started)
			f := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
			perDay[f] += w.TimeSpentSeconds
		}
	}
	return perDay
}
