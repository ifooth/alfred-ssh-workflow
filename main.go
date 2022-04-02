package main

// Package is called aw
import (
	"fmt"

	aw "github.com/deanishe/awgo"

	"github.com/ifooth/alfred-ssh-workflow/config"
)

// Workflow is the main API
var (
	wf  *aw.Workflow
	cfg *aw.Config
)

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
	cfg = aw.NewConfig()
}

// Your workflow starts here
func run() {
	conf_path := cfg.GetString("config")
	conf, err := config.ReadConfig(conf_path)
	if err != nil {
		panic(err)
	}

	for _, s := range conf.SSHConfigs {
		fmt.Println("lei1")
		if err := s.HandleItem(wf); err != nil {
			fmt.Println("lei2")
			panic(err)
		}
	}

	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
