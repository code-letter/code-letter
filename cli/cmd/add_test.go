package cmd

import (
	"github.com/redxiiikk/code-letter-cli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func Test_parseLabels(t *testing.T) {
	t.Run("should success parse", func(t *testing.T) {
		type args struct {
			labels []string
		}
		tests := []struct {
			name string
			args args
			want map[string]string
		}{
			{
				name: "given null string arrays when parse should get empty labels",
				args: args{labels: nil},
				want: map[string]string{},
			},
			{
				name: "given empty string arrays when parse should get empty labels",
				args: args{labels: []string{}},
				want: map[string]string{},
			},
			{
				name: "given label string arrays when parse should get empty labels",
				args: args{labels: []string{"name1:value1", "name2:value2"}},
				want: map[string]string{
					"name1": "value1",
					"name2": "value2",
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := parseLabels(tt.args.labels); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("parseLabels() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("given wrong string arrays when parse should exit program", func(t *testing.T) {
		if os.Getenv("TEST_SUM_COMMAND_Test_parseLabels") == "true" {
			parseLabels([]string{"test-test1"})
		} else {
			cmd := exec.Command(os.Args[0], "-test.run=Test_parseLabels")
			cmd.Env = append(os.Environ(), "TEST_SUM_COMMAND_Test_parseLabels=true")
			err := cmd.Run()

			if e, ok := err.(*exec.ExitError); ok && !e.Success() {
				return
			}
			t.Fatalf("process ran with err %v, want exit status 1", err)
		}
	})
}

var emptyFun = func(cmd *cobra.Command, args []string) {

}

func Test_AddCommand(t *testing.T) {
	beforeEach := func() *cobra.Command {
		addCommand := AddCommand()

		return addCommand
	}

	validParseLabels := func(t *testing.T, flags *pflag.FlagSet) map[string]string {
		if flags == nil {
			t.Fatalf("can't get flags when execute command")
		}

		labelStrings, err := flags.GetStringArray(addCmdLabelsFlagName)
		if err != nil {
			t.Fatalf("can't get label flag value")
		}

		if len(labelStrings) != 1 {
			t.Fatalf("get another value when given empty lables")
		}

		labels := parseLabels(labelStrings)
		return labels
	}

	validDefaultLabelValue := func(t *testing.T, labels map[string]string) {
		timestamp, isExisted := labels[defaultLabelTimestampName]
		if !isExisted {
			t.Fatalf("can't contain default label: %v", labels)
		}

		actualNowTimestampString, err := strconv.Atoi(timestamp)
		if err != nil {
			t.Fatalf("can't prase defalut timestamp label value")
		}
		actualNowTime := time.Unix(int64(actualNowTimestampString), 0)

		if actualNowTime.Sub(time.Now()).Milliseconds() > 100 {
			t.Fatalf("can't get now timestamp string")
		}
	}

	t.Run("given args that count is wrong when execute should get error", func(t *testing.T) {
		//given
		command := beforeEach()
		command.Run = emptyFun
		command.SetArgs([]string{"arg1", "arg1", "arg1"})

		// when
		err := command.Execute()

		// then
		if err == nil {
			t.Fatalf("can't valida args count")
		}
	})

	t.Run("given args when execute should not error and execute run function", func(t *testing.T) {
		//given
		command := beforeEach()
		isCall := false
		command.Run = func(cmd *cobra.Command, args []string) {
			isCall = true
		}
		command.SetArgs([]string{"arg1", "arg1", "arg1", "arg1"})

		// when
		err := command.Execute()

		// then
		if err != nil || !isCall {
			t.Fatalf("can't execute run function given right args")
		}
	})

	t.Run("given empty flags when execute should just get default labels", func(t *testing.T) {
		//given
		command := beforeEach()
		isCall := false
		var flags *pflag.FlagSet
		command.Run = func(cmd *cobra.Command, args []string) {
			isCall = true
			flags = cmd.Flags()

		}
		command.SetArgs([]string{"arg1", "arg1", "arg1", "arg1"})

		// when
		err := command.Execute()

		// then
		if err != nil || !isCall {
			t.Fatalf("can't execute run function given right args")
		}

		labels := validParseLabels(t, flags)
		if len(labels) != 1 {
			t.Fatalf("have too many label by prased flag: %v", labels)
		}
		validDefaultLabelValue(t, labels)
	})

	t.Run("given another flag when execute should get error that can't set not register flag", func(t *testing.T) {
		//given
		command := beforeEach()
		command.Run = func(cmd *cobra.Command, args []string) {
		}

		command.SetArgs([]string{"arg1", "arg1", "arg1", "arg1"})

		// when
		err := command.Flags().Set("another flag name", "another flag value")

		// then
		if err == nil {
			t.Fatalf("can't get error when set another flag value: %v", err)
		}
	})
}

type mockService struct {
	isCallPersist int
	isCallReadAll int
	readAllResult []internal.Comment
}

func (service *mockService) Persist() {
	service.isCallPersist += 1
}
func (service *mockService) ReadAll() (result []internal.Comment) {
	service.isCallReadAll += 1
	return service.readAllResult
}

func Test_run(t *testing.T) {
	t.Run("given right line number string args when run should parse this arg", func(t *testing.T) {
		factory := internal.NewServiceFactory()
		err := factory.RegisterService("mock-service", func(config *internal.Config, comment *internal.Comment) internal.Service {
			return &mockService{}
		})
		if err != nil {
			t.Fatalf("can't register mock service create function in factory")
		}

		// TODO: 完善测试
	})
}
