package tools

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	Error(fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

//goland:noinspection GoUnusedExportedFunction
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m[INFO] %s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	fmt.Printf("\x1b[31;1m[ERROR] %s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func Debug(format string, args ...interface{}) {
	fmt.Printf("\u001B[36;1m[DEBUG] %s\u001B[0m\n", fmt.Sprintf(format, args...))
}

func NowTimestampString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func IsBlankString(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func StrArrayToIntArray(array []string) (result []int, err error) {
	for _, i := range array {
		number, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}

		result = append(result, number)
	}

	return
}
