package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge [source branch] [target branch]",
	Short: "Merge source branch into target branch",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sourceBranch := args[0]
		targetBranch := args[1]

		// Checkout the target branch
		if err := runCommand("git", "checkout", targetBranch); err != nil {
			fmt.Println(err)
			return
		}

		// Pull the latest changes
		if err := runCommand("git", "pull"); err != nil {
			fmt.Println(err)
			return
		}

		// Merge the source branch into the target branch
		if err := runCommand("git", "merge", sourceBranch); err != nil {
			fmt.Println(err)
			return
		}

		// Push the changes to the remote repository
		if err := runCommand("git", "push"); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Successfully merged %s into %s\n", sourceBranch, targetBranch)
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
