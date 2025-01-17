package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute git merge recursively",
	Run: func(cmd *cobra.Command, args []string) {
		err := mergeRecersively(args)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
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

func mergeRecersively(branches []string) error {
	for i := 0; i < len(branches)-1; i++ {
		err := merge(branches[i], branches[i+1])
		if err != nil {
			return err
		}
	}
	return nil
}

func merge(srcBranch string, targetBranch string) error {
	if err := runCommand("git", "checkout", targetBranch); err != nil {
		return err
	}

	if err := runCommand("git", "pull"); err != nil {
		return err
	}

	if err := runCommand("git", "merge", "remotes/origin/"+srcBranch, "--no-ff"); err != nil {
		return err
	}

	if err := runCommand("git", "push"); err != nil {
		return err
	}

	fmt.Printf("Successfully merged [%s] into [%s]\n", srcBranch, targetBranch)
	return nil
}
