package cmd

import (
	"log"

	"github.com/jkamenik/log-pruner/scanner"
	"github.com/jkamenik/log-pruner/util"
	"github.com/spf13/cobra"
)

var dryRun = false

// pruneCmd represents the prune command
var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Prunes files and exits",
	Long: `prune scans all provided paths, deletes the files that need to be deleted and then exits.

It will only exit with an error if there is a problem reading or deleting the files.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, path := range allSettings.paths {
			err := prune(path, allSettings.fileAgeMax, util.BytesFromGb(allSettings.pathTargetSize), dryRun)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func prune(directory string, fileAge int, maxSize int64, dryRun bool) error {
	absPath, err := util.AbsPath(directory)
	if err != nil {
		return err
	}

	path, err := scanner.NewPath(absPath, fileAge, maxSize)
	if err != nil {
		return err
	}

	path.MarkOldFiles()
	if !dryRun {
		err = path.Prune()
		if err != nil {
			return err
		}
	}

	path.MarkFileUntilFit()
	if !dryRun {
		err = path.Prune()
		if err != nil {
			return err
		}
	} else {
		log.Printf("Would have pruned the following:")
		for _, file := range path.Files() {
			if file.WillPrune() {
				log.Printf("  %s (%dbytes)", file.AbsPath(), file.Size())
			}
		}
		log.Printf("Would have saved %fGb", util.GbFromBytes(path.TotalSize()-path.TotalAfterPrune()))
	}

	return nil
}

func init() {
	rootCmd.AddCommand(pruneCmd)

	pruneCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Dry-run to show the output of what would happen without actually performing the work.")
}
