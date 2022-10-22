package internal

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type reviewCommentStatus string

const (
	OPEN reviewCommentStatus = "OPEN"
)

func (status reviewCommentStatus) string() string {
	return string(status)
}

type codeReviewComment struct {
	repo   *git.Repository
	config *reviewCommentConfig

	CommitHash plumbing.Hash     `json:"commitHash"`
	FilePath   string            `json:"filePath"`
	Lines      []int             `json:"lines"`
	Comment    string            `json:"comment"`
	Labels     map[string]string `json:"labels"`
}

type reviewCommentConfig struct {
	path           string
	commentDirPath string
}
