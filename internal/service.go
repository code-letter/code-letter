package internal

import (
	"github.com/go-git/go-git/v5"
	"github.com/redxiiikk/code-review-tool-cli/tools"
)

func AddCodeReviewComment(
	repoPath string,
	commitHashStr string,
	filePath string,
	lines []int,
	reviewComment string,
	labels map[string]string,
) {
	config := newConfig(repoPath)
	comment := newReviewComment(openRepo(repoPath), config, commitHashStr, filePath, lines, reviewComment, labels)

	comment.valid()
	comment.persist()
}

func openRepo(repoPath string) *git.Repository {
	repo, err := git.PlainOpen(repoPath)
	tools.CheckIfError(err)

	return repo
}
