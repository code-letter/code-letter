package internal

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/redxiiikk/code-letter-cli/tools"
	"os"
	"path"
)

func newReviewComment(
	repo *git.Repository,
	config *reviewCommentConfig,
	commitHashStr string,
	filePath string,
	lines []int,
	reviewComment string,
	labels map[string]string,
) *codeReviewComment {
	commitHash := plumbing.NewHash(commitHashStr)
	commitObject, err := repo.CommitObject(commitHash)
	if err == plumbing.ErrObjectNotFound {
		tools.CheckIfError(err)
	}
	addDefaultLabel(labels, commitObject.Author.Name, commitObject.Author.Email)

	comment := codeReviewComment{
		repo:       repo,
		config:     config,
		commitHash: commitHash,

		CommitHashStr: commitHashStr,
		FilePath:      filePath,
		Lines:         lines,
		Comment:       reviewComment,
		Labels:        labels,
	}
	return &comment
}

func addDefaultLabel(labels map[string]string, author, email string) {
	labels["commitAuthor"] = author
	labels["commitEmail"] = email

	if _, isExisted := labels["reviewAuthor"]; !isExisted {
		labels["reviewAuthor"] = author
	}

	if _, isExisted := labels["reviewEmail"]; !isExisted {
		labels["reviewEmail"] = email
	}

	if _, isExisted := labels["status"]; !isExisted {
		labels["status"] = OPEN.string()
	}
}

func (comment *codeReviewComment) valid() {
	var errorMessages []string

	commitObject, err := comment.repo.CommitObject(comment.commitHash)
	if err == plumbing.ErrObjectNotFound {
		errorMessages = append(errorMessages, "commit object not found by hash: "+comment.CommitHashStr)
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
	comment.config.createCommentDirWhenNotExisted()

	reviewCommentHistory := comment.readReviewCommentHistory()
	reviewCommentHistory = append(reviewCommentHistory, *comment)

	bytes, err := json.Marshal(reviewCommentHistory)
	tools.CheckIfError(err)

	err = os.WriteFile(comment.storePath(), bytes, 0644)
	tools.CheckIfError(err)
}

func (comment *codeReviewComment) readReviewCommentHistory() (result []codeReviewComment) {
	result = []codeReviewComment{}

	reviewCommentStorePath := comment.storePath()

	_, err := os.Stat(reviewCommentStorePath)

	if err != nil {
		if !os.IsNotExist(err) {
			tools.CheckIfError(err)
		}
	} else {
		file, err := os.ReadFile(reviewCommentStorePath)
		tools.CheckIfError(err)

		err = json.Unmarshal(file, &result)
		tools.CheckIfError(err)
	}

	return
}

func (comment *codeReviewComment) storePath() string {
	reviewCommentFileName := comment.Labels["reviewAuthor"] + ":" + comment.Labels["reviewEmail"]
	return path.Join(comment.config.commentDirPath, reviewCommentFileName)
}
