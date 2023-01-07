package cmd

import (
	"github.com/redxiiikk/code-letter-cli/internal"
	"github.com/redxiiikk/code-letter-cli/tools"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

const (
	addCmdName           = "add"
	addCmdShortDescriber = "add code review comments"
	addCmdLongDescriber  = `add code review comments to specified file
arguments:
  1. commit hash
  2. blob object hash
  3. review comment`

	addCmdLabelsFlagName = "label"
)

var labels []string

var addCmd = &cobra.Command{
	Use:   addCmdName,
	Short: addCmdShortDescriber,
	Long:  addCmdLongDescriber,
	Args:  cobra.ExactArgs(4),
	Run:   run,
}

func init() {
	initFlag()
}

func initFlag() {
	var defaultLabelsFlagValues = []string{"timestamp:" + tools.NowTimestampString()}
	addCmd.Flags().StringArrayVar(&labels, addCmdLabelsFlagName, defaultLabelsFlagValues, "add label for code review comments")
}

func run(cmd *cobra.Command, args []string) {
	commitHash := args[0]
	filePath := args[1]

	lineNumbers, err := tools.StrArrayToIntArray(strings.Split(args[2], ","))
	tools.CheckIfError(err)

	reviewComment := args[3]

	internal.AddCodeReviewComment(parseRepoPath(cmd), commitHash, filePath, lineNumbers, reviewComment, parseLabels(labels))
}

func parseLabels(labels []string) map[string]string {
	if labels == nil || len(labels) == 0 {
		return nil
	}

	result := make(map[string]string, len(labels))

	for _, label := range labels {
		labelArray := strings.Split(label, ":")

		if len(labelArray) != 2 {
			tools.Error("label is invalid: " + label)
			os.Exit(1)
		}

		result[labelArray[0]] = labelArray[1]
	}

	return result
}
