package internal

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type codeReviewComment struct {
	repo       *git.Repository
	CommitHash plumbing.Hash     `json:"commitHash"`
	FilePath   string            `json:"filePath"`
	Lines      []int             `json:"lines"`
	Comment    string            `json:"comment"`
	Labels     map[string]string `json:"labels"`
}
