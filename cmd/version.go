package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mergectl",
	Run: func(cmd *cobra.Command, args []string) {
		version := "1.0.0"
		fmt.Println("mergectl version:", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
