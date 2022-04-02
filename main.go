package main

// Package is called aw
import (
	"fmt"

	aw "github.com/deanishe/awgo"
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
	// Add a "Script Filter" result
	wf.NewItem("First result!")

	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	config := cfg.GetString("config")
	fmt.Println("lei", config)
	wf.Run(run)
}
