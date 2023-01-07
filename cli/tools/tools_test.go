package tools

import (
	"errors"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestIsBlankString(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "should get true when param is empty string",
			args: "",
			want: true,
		},
		{
			name: "should get true when param is blank string",
			args: "        ",
			want: true,
		},
		{
			name: "should get false when param is not a blank",
			args: "THIS IS NOT A BLANK STRING",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlankString(tt.args); got != tt.want {
				t.Errorf("IsBlankString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StrArrayToIntArray(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantResult []int
		wantErr    bool
	}{
		{
			name:       "should get int array given a element is number string in array",
			args:       []string{"1", "2"},
			wantResult: []int{1, 2},
			wantErr:    false,
		},
		{
			name:    "should get error given a element is float string in array",
			args:    []string{"1.2", "2.3"},
			wantErr: true,
		},
		{
			name:    "should get error given a element is alphabet in array",
			args:    []string{"A", "B"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := StrArrayToIntArray(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrArrayToIntArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("StrArrayToIntArray() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNowTimestampString(t *testing.T) {
	expectedNowTime := time.Now()

	actualNowTimestampString, err := strconv.Atoi(NowTimestampString())
	if err != nil {
		t.Errorf("can't convert string about now timestamp to int")
		t.Fail()
	}

	actualNowTime := time.Unix(int64(actualNowTimestampString), 0)

	if actualNowTime.Sub(expectedNowTime).Milliseconds() > 100 {
		t.Errorf("can't get now timestamp string")
		t.Fail()
	}
}

func TestCheckIfError(t *testing.T) {
	t.Run("should normal exist when don't given error", func(t *testing.T) {
		if os.Getenv("TEST_SUM_COMMAND_TestCheckIfError_1") == "true" {
			CheckIfError(nil)
			return
		} else {
			cmd := exec.Command(os.Args[0], "-test.run=TestCheckIfError")
			cmd.Env = append(os.Environ(), "TEST_SUM_COMMAND_TestCheckIfError_1=true")

			err := cmd.Run()

			if e, ok := err.(*exec.ExitError); ok && !e.Success() {
				t.Fatalf("process ran with err %v, want exit status 1", err)
			}
		}
	})

	t.Run("should abnormal exist and exist code is 1 given error", func(t *testing.T) {
		if os.Getenv("TEST_SUM_COMMAND_TestCheckIfError_2") == "true" {
			CheckIfError(errors.New("THIS IS A ERROR FOR TEST CheckIfError"))
			return
		} else {
			cmd := exec.Command(os.Args[0], "-test.run=TestCheckIfError")
			cmd.Env = append(os.Environ(), "TEST_SUM_COMMAND_TestCheckIfError_2=true")

			err := cmd.Run()
			if e, ok := err.(*exec.ExitError); ok && !e.Success() {
				return
			}
			t.Fatalf("process ran with err %v, want exit status 1", err)
		}
	})
}
