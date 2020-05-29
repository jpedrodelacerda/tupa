package cmd

import (
	"github.com/jpedrodelacerda/tupa/pkg/ignore"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a new project",
	Long: `Creates a new project. You can specify the license
and items to be ignored on the .gitignore.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createDirectory(args[0])

		ignore.CreateGitIgnore(params, args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}

func createDirectory(path string) error {
	err := os.Mkdir(path, os.FileMode(0755))
	return err
}
