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

func AddCodeReviewComment(repoPath string, commitHash plumbing.Hash, filePath string, lines []int, reviewComment string, labels map[string]string) {
	repo := openRepo(repoPath)

	comment := codeReviewComment{repo, commitHash, filePath, lines, reviewComment, labels}

	comment.valid()
	comment.persist()
}

func openRepo(repoPath string) *git.Repository {
	repo, err := git.PlainOpen(repoPath)
	tools.CheckIfError(err)

	return repo
}

func (comment *codeReviewComment) valid() {
	var errorMessages []string

	commitObject, err := comment.repo.CommitObject(comment.CommitHash)
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
