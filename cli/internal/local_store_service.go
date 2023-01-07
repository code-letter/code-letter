package internal

import (
	"encoding/json"
	"github.com/redxiiikk/code-letter-cli/tools"
	"os"
	"path"
)

const (
	defaultCommentDirPath = ".code-reviews/comments"
)

type LocalStoreService struct {
	config  *Config
	comment *Comment
}

func newLocalStoreService(config *Config, comment *Comment) Service {
	return &LocalStoreService{config, comment}
}

func (service *LocalStoreService) Persist() {
	service.createCommentDirWhenNotExisted()

	service.comment.valid()
	service.persist()
}

func (service *LocalStoreService) createCommentDirWhenNotExisted() {
	tools.Debug("create comment dir when not existed: " + service.allCommentDirPath())
	stat, err := os.Stat(service.allCommentDirPath())

	if os.IsNotExist(err) || !stat.IsDir() {
		err := os.MkdirAll(service.allCommentDirPath(), os.ModePerm)
		tools.CheckIfError(err)
	}
}

func (service *LocalStoreService) persist() {
	reviewCommentHistory := service.ReadAll()
	reviewCommentHistory = append(reviewCommentHistory, *service.comment)

	bytes, err := json.Marshal(reviewCommentHistory)
	tools.CheckIfError(err)

	err = os.WriteFile(service.currentCommentFilePath(), bytes, 0644)
	tools.CheckIfError(err)
}

func (service *LocalStoreService) allCommentDirPath() string {
	return path.Join(service.config.localRepoPath, defaultCommentDirPath)
}

func (service *LocalStoreService) currentCommentFilePath() string {
	reviewCommentFileName := service.comment.Labels["reviewAuthor"] + ":" + service.comment.Labels["reviewEmail"]
	return path.Join(service.allCommentDirPath(), reviewCommentFileName)
}

func (service *LocalStoreService) ReadAll() (result []Comment) {
	result = []Comment{}

	reviewCommentStorePath := service.currentCommentFilePath()

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
