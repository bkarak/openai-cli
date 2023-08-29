package cmd_actions

import (
	"fmt"
	"github.com/spf13/cobra"
	"openai-cli/integrations/openai"
)

var modelCommand = &cobra.Command{
	Use:   "models-list",
	Short: "Lists all available models",
	RunE: func(cmd *cobra.Command, args []string) error {
		models := openai.ListModels()
		fmt.Printf("Executing: listing all models\n")

		for c := 0; c < len(models); c++ {
			fmt.Printf("Model => id: %s\n", models[c].Id)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(modelCommand)
}
