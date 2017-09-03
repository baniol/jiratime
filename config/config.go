package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"

	"gopkg.in/yaml.v2"
)

// Config struct represents the jiratime configuration structure
type Config struct {
	JiraUser         string `yaml:"jirauser"`
	JiraURL          string `yaml:"jiraurl"`
	JiraPassword     string `yaml:"jirapassword"`
	DateFrom         string `yaml:"datefrom"`
	HoursDaily       int    `yaml:"hoursdaily"`
	MaxSearchResults int    `yaml:"maxsearchresults"`
}

// @TODO pointer receiver for Config ?
func checkConfig(conf Config) error {

	configErrors := make([]string, 0)

	if conf.JiraURL == "" {
		configErrors = append(configErrors, "JiraURL")
	}
	if conf.JiraUser == "" {
		configErrors = append(configErrors, "JiraUser")
	}
	if conf.JiraPassword == "" {
		configErrors = append(configErrors, "JiraPassword")
	}
	if conf.DateFrom == "" {
		configErrors = append(configErrors, "DateFrom")
	}
	if conf.HoursDaily == 0 {
		configErrors = append(configErrors, "HoursDaily")
	}
	if conf.MaxSearchResults == 0 {
		configErrors = append(configErrors, "MaxSearchResults")
	}

	if len(configErrors) > 0 {
		errMsg := strings.Join(configErrors[:], ", ")
		return fmt.Errorf("%s not configured", errMsg)
	}

	return nil
}

func readFile(cfgFile string) ([]byte, error) {
	source, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	return source, nil
}

func readConfig(fileContent []byte) (Config, error) {
	config := Config{}
	err := yaml.Unmarshal(fileContent, &config)
	return config, err
}

// getConfigFilePath returns a path to jiratime config file
func getConfigFilePath(pathParam string) (string, error) {

	// Path specified as cli param
	if pathParam != "" {
		return pathParam, nil
	}

	userHome, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%s/jiratimeconfig.yml", userHome)
	return fileName, nil
}

// GetConfig returns the Config instance
// pathParam is an optional param passed from user input
func GetConfig(pathParam string) (Config, error) {

	filePath, _ := getConfigFilePath(pathParam)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return Config{}, err
	}
	content, err := readFile(filePath)
	if err != nil {
		return Config{}, err
	}

	conf, err := readConfig(content)

	if err != nil {
		return Config{}, err
	}

	err = checkConfig(conf)

	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
