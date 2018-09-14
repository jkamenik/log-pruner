package cmd

import (
	"log"
	"time"

	"github.com/jkamenik/log-pruner/util"
	"github.com/spf13/cobra"
)

var wait int

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Like prune but runs forever",
	Long:  `Continually scans and prunes all the paths provided.  When done it sleeps for a time and then starts again.`,
	Run: func(cmd *cobra.Command, args []string) {
		for true {
			for _, path := range allSettings.paths {
				err := prune(path, allSettings.fileAgeMax, util.BytesFromGb(allSettings.pathTargetSize), allSettings.dryRun)
				if err != nil {
					log.Fatal(err)
				}
			}

			log.Printf("Done Pruning, waiting %ds for next prune cycle", wait)
			time.Sleep(time.Duration(wait) * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)

	monitorCmd.Flags().IntVar(&wait, "wait", 300, "Wait time in seconds between successive checks.  Note: all directories are scanned and pruned entirely before the wait is triggered.")
}
