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

const labelNameValueSeparator = ":"
const defaultLabelTimestampName = "timestamp"

const lineNumberSeparator = ","

func AddCommand() *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   addCmdName,
		Short: addCmdShortDescriber,
		Long:  addCmdLongDescriber,
		Args:  cobra.ExactArgs(4),
		Run:   run,
	}

	initFlag(addCmd)

	return addCmd
}

func initFlag(command *cobra.Command) {
	var defaultLabelsFlagValues = []string{defaultLabelTimestampName + labelNameValueSeparator + tools.NowTimestampString()}
	command.Flags().StringArray(addCmdLabelsFlagName, defaultLabelsFlagValues, "add label for code review comments")
}

func run(cmd *cobra.Command, args []string) {
	var (
		commitHash            = args[0]
		filePath              = args[1]
		lineNumberArrayString = args[2]
		reviewComment         = args[3]
	)

	lineNumbers, err := tools.StrArrayToIntArray(strings.Split(lineNumberArrayString, lineNumberSeparator))
	tools.CheckIfError(err)

	labelStrings, err := cmd.Flags().GetStringArray(addCmdLabelsFlagName)
	tools.CheckIfError(err)
	labels := parseLabels(labelStrings)

	localRepoPath := parseRepoPath(cmd)

	config := internal.NewConfig(localRepoPath)
	comment := internal.NewComment(localRepoPath, commitHash, filePath, lineNumbers, reviewComment, labels)

	serviceFactor := internal.NewServiceFactory()
	service, err := serviceFactor.CreateService(config, comment)
	tools.CheckIfError(err)

	service.Persist()
}

func parseLabels(labels []string) map[string]string {
	if labels == nil || len(labels) == 0 {
		return map[string]string{}
	}

	result := make(map[string]string, len(labels))

	for _, label := range labels {
		labelArray := strings.Split(label, labelNameValueSeparator)

		if len(labelArray) != 2 {
			tools.Error("label is invalid: " + label)
			os.Exit(1)
		}

		result[labelArray[0]] = labelArray[1]
	}

	return result
}
