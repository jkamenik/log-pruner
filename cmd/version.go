package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Symantic version that is used both in the version command and the build scripts.
var version = "1.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of the application.",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:
	//
	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
