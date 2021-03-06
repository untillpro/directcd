package main

import (
	"github.com/spf13/cobra"
	gc "github.com/untillpro/gochips"
)

func main() {

	var rootCmd = &cobra.Command{Use: "directcd"}
	rootCmd.PersistentFlags().BoolVarP(&gc.IsVerbose, "verbose", "v", false, "Verbose output")

	// cmdPull
	{
		var cmdPull = &cobra.Command{
			Use:   "pull --repo <main-repo> --replace <repo2=<repo-to-replace]> [args]",
			Short: "Pull sources from given repos",
			Run:   runCmdPull,
		}

		cmdPull.PersistentFlags().StringVarP(&binaryName, "output", "o", "", "Output binary name")
		cmdPull.PersistentFlags().StringVarP(&workingDir, "working-dir", "w", ".", "Working directory")
		cmdPull.PersistentFlags().Int32VarP(&timeoutSec, "timeout", "t", 10, "Timeout")
		cmdPull.PersistentFlags().StringVarP(&mainRepo, "repo", "r", "", "Main repository")
		cmdPull.PersistentFlags().StringSliceVar(&argReplaced, "replace", []string{}, "Repositories to be replaced")
		cmdPull.MarkPersistentFlagRequired("output")
		cmdPull.MarkPersistentFlagRequired("repo")

		rootCmd.AddCommand(cmdPull)
	}

	rootCmd.Execute()

}
