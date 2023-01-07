package internal

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/redxiiikk/code-letter-cli/tools"
	"os"
)

type Comment struct {
	repo       *git.Repository
	commitHash plumbing.Hash

	CommitHashStr string            `json:"commitHash"`
	FilePath      string            `json:"filePath"`
	Lines         []int             `json:"lines"`
	Comment       string            `json:"comment"`
	Labels        map[string]string `json:"labels"`
}

func NewComment(
	localRepoPath string,
	commitHashStr string,
	filePath string,
	lines []int,
	reviewComment string,
	labels map[string]string,
) *Comment {
	repo := openRepo(localRepoPath)
	commitHash := plumbing.NewHash(commitHashStr)
	commitObject, err := repo.CommitObject(commitHash)
	if err == plumbing.ErrObjectNotFound {
		tools.CheckIfError(err)
	}

	addDefaultLabel(labels, commitObject.Author.Name, commitObject.Author.Email)

	comment := Comment{
		commitHash: commitHash,

		CommitHashStr: commitHashStr,
		FilePath:      filePath,
		Lines:         lines,
		Comment:       reviewComment,
		Labels:        labels,
	}
	return &comment
}

func openRepo(repoPath string) *git.Repository {
	repo, err := git.PlainOpen(repoPath)
	tools.CheckIfError(err)

	return repo
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

func (comment *Comment) valid() {
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
