package assert

import (
	"fmt"
	"testing"
)

var resetText = "\033[0m"
var startRedText = "\033[31m"
var startGreenText = "\033[32m"

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if actual != expected {
		testFailed(t, "\n\tACTUAL: %v\n\tEXPECTED: %v", actual, expected)
	} else {
		testPassed(t, "")
	}
}

func testFailed(t *testing.T, format string, message ...any) {
	output := fmt.Sprintf(format, message...)
	output = "\nTEST FAILED" + output
	t.Errorf(startRedText + output + resetText)
}

func testPassed(t *testing.T, format string, message ...any) {
	output := fmt.Sprintf(format, message...)
	output = "\nTEST PASSED" + output
	t.Logf(startGreenText + output + resetText)
}
