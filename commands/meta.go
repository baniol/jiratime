package commands

import (
	"flag"
	"fmt"

	"github.com/baniol/jiratime/config"
	"github.com/mitchellh/cli"
)

type Meta struct {
	Ui cli.Ui
}

func (m *Meta) Config(path string) (*config.Config, error) {

	conf, err := config.GetConfig(path)
	if err != nil {
		m.Ui.Error(fmt.Sprintf("Error: %s", err))
		return nil, err
	}
	return &conf, nil
}

func (m *Meta) FlagSet(n string) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)
	var configPath string
	f.StringVar(&configPath, "config", "", "Configuration file")
	return f
}
