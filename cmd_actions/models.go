package cmd_actions

import (
	"fmt"
	"github.com/spf13/cobra"
	"openai-cli/integrations/openai"
)

var modelCmd = &cobra.Command{
	Use: "models",
	RunE: func(cmd *cobra.Command, args []string) error {
		for c := 0; c < len(args); c++ {
			print(args[c])
		}

		models := openai.ListModels()

		fmt.Printf("Executing: listing all models from openai\n")

		for c := 0; c < len(models); c++ {
			fmt.Printf("Model => id: %s\n", models[c].Id)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
