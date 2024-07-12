package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mergectl",
	Short: "USAGE: mergectl [source branch] [target branch]",
	Long:  `mergectl is a tool of "git merge".`,
	Args:  cobra.ExactArgs(2), // 複数の引数を受け入れて、再帰的に0, 1マージをするようにしないといけない
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

		if err := runCommand("git", "merge", "remotes/origin/"+sourceBranch, "--no-ff"); err != nil {
			fmt.Println(err)
			return
		}

		if err := runCommand("git", "push"); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Successfully merged [%s] into [%s]\n", sourceBranch, targetBranch)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runCommand(name string, args ...string) error {
	command := name
	for _, a := range args {
		command += " " + a
	}
	fmt.Printf("Running... [%s]\n", command)

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
