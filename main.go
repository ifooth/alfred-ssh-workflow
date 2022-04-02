package main

import (
	aw "github.com/deanishe/awgo"

	"github.com/ifooth/alfred-ssh-workflow/config"
)

const (
	defaultConfig = "~/.config/sshmgr/config.yml"
)

var (
	wf  *aw.Workflow
	cfg *aw.Config
)

func init() {
	wf = aw.New()
	cfg = aw.NewConfig()
}

func run() {
	conf, err := config.ReadConfig(cfg.GetString("config", defaultConfig))
	if err != nil {
		panic(err)
	}

	for _, s := range conf.SSHConfigs {
		if err := s.HandleItem(wf); err != nil {
			panic(err)
		}
	}

	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
