package cmd

import (
	"fmt"
	"os"

	"github.com/jpedrodelacerda/tupa/pkg/ignore"
	"github.com/jpedrodelacerda/tupa/pkg/license"

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

		ignore.CreateGitIgnore(ignoreParams, args[0])
		license.CreateLicense(licenseName, args[0], author)
		fmt.Printf("Project created!\nAuthor:\t\t%s\nLicense:\t%s\n", author, licenseName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}

func createDirectory(path string) error {
	err := os.Mkdir(path, os.FileMode(0755))
	return err
}
