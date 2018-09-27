package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type settings struct {
	paths          []string
	pathTargetSize int
	fileAgeMax     int
	dryRun         bool
}

var allSettings = settings{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "1.0.0",
	Use:     "log-pruner",
	Long: `A command that either scan or monitors a diretory and removes them based on age or size criteria.

All supplied directories are scanned, and then pruned of files older then their max age.  Then the files are rescanned and pruned if the path is still too large.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceVar(&allSettings.paths, "path", []string{"/logs"}, "One or more paths for the command to check.  Both comma separated and multiple arguments are allowed.  Each path is taken separately in its calculations.")
	rootCmd.PersistentFlags().IntVar(&allSettings.pathTargetSize, "path-target-size", 10, "The target size for the entire path in GB.  Files will be deleted if the target size is greater then this value, even if no files are older then the max age.")
	rootCmd.PersistentFlags().IntVar(&allSettings.fileAgeMax, "file-max-age", 3, "The max number of days to keep a file before it is deleted.")
	rootCmd.PersistentFlags().BoolVar(&allSettings.dryRun, "dry-run", false, "Dry-run to show the output of what would happen without actually performing the work.")
}
