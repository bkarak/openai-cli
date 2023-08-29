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
	// intialize the createImageCommand
	rootCmd.AddCommand(createImageCommand)
	createImageCommand.Flags().String("prompt", "", "the description of the new image")
	createImageCommand.Flags().Int("cardinality", 1, "how many images should be generated (1 to 10), default is 1")
	createImageCommand.Flags().String("size", "256x256", "size of the image, one of 256x256, 512x512, 1024x1024")

	rootCmd.AddCommand(createImageEditCommand)
	rootCmd.AddCommand(createImageVariationCommand)
}
