package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mergectl",
	Short: "USAGE: mergectl [source branch] [target branch]",
	Long:  `mergectl is a tool of "git merge".`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sourceBranch := args[0]
		targetBranch := args[1]

		if err := runCommand("git", "checkout", targetBranch); err != nil {
			fmt.Println(err)
			return
		}

		if err := runCommand("git", "pull"); err != nil {
			fmt.Println(err)
			return
		}

		if err := runCommand("git", "merge", sourceBranch); err != nil {
			fmt.Println(err)
			return
		}

		if err := runCommand("git", "push"); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Successfully merged %s into %s\n", sourceBranch, targetBranch)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}