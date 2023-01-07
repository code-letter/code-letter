package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/redxiiikk/code-letter-cli/tools"
	"github.com/spf13/cobra"
	"os"
)

const (
	rootCmdShortDescriber = "record code review suggestion in local file"
	rootCmdLongDescriber  = "a tools to help record code review suggestions to store local file"

	rootCmdRepoFlagName = "repo"
)

var rootCmd = &cobra.Command{
	Short: rootCmdShortDescriber,
	Long:  rootCmdLongDescriber,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		checkRepoFlag(cmd)
	},
}

func init() {
	rootCmd.AddCommand(AddCommand())
	rootCmd.PersistentFlags().String(rootCmdRepoFlagName, "", "repo path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkRepoFlag(cmd *cobra.Command) {
	realRepoPath := parseRepoPath(cmd)

	_, err := git.PlainOpen(realRepoPath)
	if err == git.ErrRepositoryNotExists {
		tools.Error("specified path is not a code repository, please check repo path again!!!")
		os.Exit(1)
	}
}

func parseRepoPath(cmd *cobra.Command) string {
	inputRepoPath, err := cmd.Flags().GetString(rootCmdRepoFlagName)
	tools.CheckIfError(err)

	if inputRepoPath != "" {
		return inputRepoPath
	} else {
		currentDirPath, err := os.Getwd()
		tools.CheckIfError(err)

		tools.Debug("use current dir path as repo: " + currentDirPath + "'")
		return currentDirPath
	}
}
