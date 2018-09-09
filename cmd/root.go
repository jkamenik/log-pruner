package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "1.0.0",
	Use:     "log-pruner",
	Long: `A command that either scan or monitors a diretory full of files and removes them based on age or size criteria.

All supplied directories are scanned, and then pruned of files larger then their max age.  Then the files are rescanned and pruned if the path is still too large.
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
	rootCmd.PersistentFlags().StringSlice("path", []string{"/logs"}, "One or more paths for the command to check.  Both comma separated and multiple arguments are allowed.  Each path is taken separately in its calculations.")
	rootCmd.PersistentFlags().Int("path-target-size", 10, "The target size for the entire path in GB.  Files will be deleted if the target size is greater then this value, even if no files are older then the max age.")
	rootCmd.PersistentFlags().Int("file-max-age", 3, "The max number of days to keep a file before it is deleted.")
}
