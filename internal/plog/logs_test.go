package plog

import (
	"path/filepath"
	"runtime"
	"testing"
)

var testLogFilePath = ""

const pathFromThisFileToLogFile = "..\\..\\..\\logs\\testing\\"
const testLogFileName = "testLogs"
const testLogFileExtension = "log"

func init() {
	var _, callingFile, _, _ = runtime.Caller(0)
	testLogFilePath = filepath.Dir(callingFile)
	testLogFilePath = filepath.Clean(testLogFilePath + pathFromThisFileToLogFile)
}

func TestLogWrapper_CliOutput(t *testing.T) {
	str := "Testing the logger CLI output."
	logger := New(testLogFilePath, testLogFileName, testLogFileExtension)

	logger.Debug(str)
	logger.Info(str)
	logger.Warn(str)
	logger.Error(str)
}
