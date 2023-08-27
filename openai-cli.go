package main

import (
	"openai-cli/cmd_actions"
	"openai-cli/system"
	"os"
)

func main() {
	if !system.CheckOpenAIConfiguration() {
		system.Fatal("missing openai configuration")
		os.Exit(1)
	}

	cmd_actions.Execute()
}
