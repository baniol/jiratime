package worklog

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	gojira "github.com/andygrunwald/go-jira"
	"github.com/baniol/jiratime/config"
)

type searchResult struct {
	Issues     []gojira.Issue `json:"issues" structs:"issues"`
	StartAt    int            `json:"startAt" structs:"startAt"`
	MaxResults int            `json:"maxResults" structs:"maxResults"`
	Total      int            `json:"total" structs:"total"`
}

type jiraFakeSession struct{}

func (s *jiraFakeSession) getUserWorklog() ([]gojira.Issue, error) {
	jsonFile := "worklog_fixture.json"
	source, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	v := new(searchResult)
	json.Unmarshal(source, &v)

	return v.Issues, nil
}

// https://stackoverflow.com/a/18208542
func eqMap(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func Test_MapHoursPerTicket(t *testing.T) {
	client := new(jiraFakeSession)

	tickets := GetUserTickets(client)

	perTicket, total := MapHoursPerTicket(tickets)
	expected := 162000
	if total != expected {
		t.Errorf("Expected %v but got %v", expected, total)
	}

	expectedMap := map[string]int{
		"TK-123": 7200,
		"TK-734": 7200,
		"TK-234": 57600,
		"TK-987": 90000,
	}

	eq := eqMap(perTicket, expectedMap)
	if !eq {
		t.Errorf("Expected %v but got %v", expectedMap, perTicket)
	}

}

func Test_MapHoursPerDay(t *testing.T) {
	client := new(jiraFakeSession)

	tickets := GetUserTickets(client)

	perDay := MapHoursPerDay(tickets)

	expectedMap := map[string]int{
		"2017-08-10": 14400,
		"2017-08-11": 25200,
		"2017-08-17": 14400,
		"2017-08-16": 25200,
		"2016-12-28": 7200,
		"2017-05-07": 7200,
		"2017-08-23": 25200,
		"2017-08-24": 25200,
		"2017-08-02": 7200,
		"2017-08-07": 10800,
	}

	eq := eqMap(perDay, expectedMap)
	if !eq {
		t.Errorf("Expected %v but got %v", expectedMap, perDay)
	}

}

func Test_NewClient(t *testing.T) {
	conf := config.Config{JiraUser: "test.user"}
	client := NewClient(&conf)
	check := reflect.TypeOf(client).String()
	expected := "*worklog.JiraSession"
	if check != expected {
		t.Errorf("Expected %v but got %v", expected, check)
	}

	// @TODO check content of jiratimeConfig here ?
}
