package internal

import (
	"github.com/redxiiikk/code-review-tool-cli/tools"
	"os"
	"path"
)

const (
	defaultConfigDir  = ".code-reviews"
	defaultCommentDir = "comments"
)

func newConfig(repoPath string) *reviewCommentConfig {
	configDirPath := path.Join(repoPath, defaultConfigDir)

	return &reviewCommentConfig{
		path:           configDirPath,
		commentDirPath: path.Join(configDirPath, defaultCommentDir),
	}
}

func (config *reviewCommentConfig) createCommentDirWhenNotExisted() {
	tools.Debug("create comment dir when not existed: " + config.commentDirPath)
	stat, err := os.Stat(config.commentDirPath)

	if os.IsNotExist(err) || !stat.IsDir() {
		err := os.MkdirAll(config.commentDirPath, os.ModePerm)
		tools.CheckIfError(err)
	}
}
