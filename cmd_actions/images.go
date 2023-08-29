package cmd_actions

import "github.com/spf13/cobra"

var createImageCommand = &cobra.Command{
	Use:   "image-create",
	Short: "Create a new image (or set of images)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var createImageEditCommand = &cobra.Command{
	Use:   "image-edit",
	Short: "Edit an image",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var createImageVariationCommand = &cobra.Command{
	Use:   "image-variation",
	Short: "Create a variation of an input image",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createImageCommand)
	rootCmd.AddCommand(createImageEditCommand)
	rootCmd.AddCommand(createImageVariationCommand)
}
