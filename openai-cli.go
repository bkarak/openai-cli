package main

import (
	"openai-cli/system"
	"os"
)

func main() {
	if !system.CheckOpenAIConfiguration() {
		system.Fatal("missing openai configuration")
		os.Exit(1)
	}
}
