package config

import (
	"testing"
)

var testPath string = "test_config.yml"

func Test_GetConfig(t *testing.T) {

	expectedUrl := "https://test.jira"
	expectedUser := "test.user"

	conf, err := GetConfig(testPath)

	if err != nil {
		t.Fatalf("Error reading config %v: ", err)
	}

	if conf.JiraURL != expectedUrl {
		t.Errorf("Expected %v but got %v", expectedUrl, conf.JiraURL)
	}

	if conf.JiraUser != expectedUser {
		t.Errorf("Expected %v but got %v", expectedUser, conf.JiraUser)
	}

}
