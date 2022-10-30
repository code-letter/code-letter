package internal

import (
	"os"
	"os/exec"
	"path"
	"reflect"
	"testing"
)

func Test_newConfig(t *testing.T) {
	type args struct {
		repoPath string
	}

	notExistedDir := "/this-is-a-test-repo-path"

	tests := []struct {
		name string
		args args
		want *reviewCommentConfig
	}{
		{
			name: "should get comment dir path given repoPath",
			args: args{
				repoPath: notExistedDir,
			},
			want: &reviewCommentConfig{
				path:           notExistedDir + "/.code-reviews",
				commentDirPath: notExistedDir + "/.code-reviews/comments",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newConfig(tt.args.repoPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reviewCommentConfig_createCommentDirWhenNotExisted(t *testing.T) {
	t.Run("should create code review dir when it's not existed", func(t *testing.T) {
		tempDir := t.TempDir()
		config := newConfig(tempDir)
		config.createCommentDirWhenNotExisted()

		if stat, err := os.Stat(tempDir + "/.code-reviews/comments"); os.IsNotExist(err) || !stat.IsDir() {
			t.Errorf("can't create comment dir")
		}
	})

	t.Run(
		"should create comment dir with same name when this existed a file exists with the same name as the folder waiting to be created",
		func(t *testing.T) {
			tempDir := t.TempDir()

			err := os.Mkdir(tempDir+"/.code-reviews", os.ModePerm)
			if err != nil {
				t.Errorf("can't create code review dir when paper test data： " + err.Error())
			}
			err = os.WriteFile(path.Join(tempDir+"/.code-reviews/comments"), []byte{}, os.ModePerm)
			if err != nil {
				t.Errorf("can't create file with same name when paper test data： " + err.Error())
			}

			if os.Getenv("TEST_SUM_COMMAND_Test_reviewCommentConfig_createCommentDirWhenNotExisted") == "true" {
				config := newConfig(tempDir)
				config.createCommentDirWhenNotExisted()

				if stat, err := os.Stat(tempDir + "/.code-reviews/comments"); os.IsNotExist(err) || !stat.IsDir() {
					t.Errorf("can't create comment dir")
				}
			} else {
				cmd := exec.Command(os.Args[0], "-test.run=Test_reviewCommentConfig_createCommentDirWhenNotExisted")
				cmd.Env = append(os.Environ(), "TEST_SUM_COMMAND_Test_reviewCommentConfig_createCommentDirWhenNotExisted=true")
				err := cmd.Run()

				if err != nil {
					if e, ok := err.(*exec.ExitError); ok != true || e.ExitCode() != 1 {
						t.Errorf("not exit when existed file that is same name with comment dir: %v, %v", ok, e)
					}
				}
			}
		},
	)
}
