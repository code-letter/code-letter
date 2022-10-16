package internal

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/redxiiikk/code-review-tool-cli/tools"
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
