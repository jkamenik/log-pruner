package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pruneCmd represents the prune command
var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Prunes files and exits",
	Long: `prune scans all provided paths, deletes the files that need to be deleted and then exits.

It will only exit with an error if there is a problem reading or deleting the files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prune called")
	},
}

func init() {
	rootCmd.AddCommand(pruneCmd)

	pruneCmd.Flags().Bool("dry-run", false, "Dry-run to show the output of what would happen without actually performing the work.")
}
