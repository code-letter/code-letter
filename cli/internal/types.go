package internal

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/redxiiikk/code-review-tool-cli/tools"
	"os"
)

type codeReviewComment struct {
	Repo       *git.Repository   `json:"-"`
	CommitHash plumbing.Hash     `json:"commitHash"`
	FilePath   string            `json:"filePath"`
	Lines      []int             `json:"lines"`
	Comment    string            `json:"comment"`
	Labels     map[string]string `json:"labels"`
}

func (comment *codeReviewComment) valid() {
	var errorMessages []string

	commitObject, err := comment.Repo.CommitObject(comment.CommitHash)
	if err == plumbing.ErrObjectNotFound {
		errorMessages = append(errorMessages, "commit object not found by hash: "+comment.CommitHash.String())
	}

	f, err := commitObject.File(comment.FilePath)
	if err == object.ErrFileNotFound {
		errorMessages = append(errorMessages, "file not found in this commit: "+comment.FilePath)
	}

	contentsByLine, err := f.Lines()
	maxLineNumber := len(contentsByLine)
	tools.CheckIfError(err)
	for _, lineNum := range comment.Lines {
		if lineNum <= 0 || lineNum > maxLineNumber {
			errorMessages = append(errorMessages, fmt.Sprintf("line number out of file: %d", lineNum))
		}
	}

	if tools.IsBlankString(comment.Comment) {
		errorMessages = append(errorMessages, "review comment is empty, please write some words")
	}

	if len(errorMessages) != 0 {
		for _, msg := range errorMessages {
			tools.Error(msg)
		}

		os.Exit(1)
	}
}

func (comment *codeReviewComment) persist() {
	bytes, err := json.Marshal(comment)
	tools.CheckIfError(err)

	fmt.Println(string(bytes))
}
