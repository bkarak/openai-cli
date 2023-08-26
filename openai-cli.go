package main

import (
	"openai-cli/integrations/openai"
	"openai-cli/system"
	"os"
)

func main() {
	if !system.CheckOpenAIConfiguration() {
		system.Fatal("missing openai configuration")
		os.Exit(1)
	}

	models := openai.ListModels()

	for c := 0; c < len(models); c++ {
		println(models[c].Id)
	}
}
