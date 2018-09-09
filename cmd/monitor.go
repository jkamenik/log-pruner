package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Like prune but runs forever",
	Long:  `Continually scans and prunes all the paths provided.  When done it sleeps for a time and then starts again.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitor called")
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)

	monitorCmd.Flags().Int("wait", 300, "Wait time in seconds between successive checks.  Note: all directories are scanned and pruned entirely before the wait is triggered.")
}
