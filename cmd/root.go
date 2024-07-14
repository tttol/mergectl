package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mergectl",
	Short: "USAGE: mergectl [source branch] [target branch]",
	Long:  `mergectl is a tool of "git merge". This command merge multiple git branches.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello mergectl!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
