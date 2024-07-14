package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mergectl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mergectl version:", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
