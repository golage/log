package log

import (
	"bytes"
	"os"
	"testing"
)

var (
	testOutput = new(bytes.Buffer)
	testLevels = []Level{LevelDebug, LevelInfo, LevelWarning, LevelError, LevelFatal}
)

func TestMain(m *testing.M) {
	resetTest()
	os.Exit(m.Run())
}

func resetTest() {
	testOutput.Reset()
	level = LevelDebug
	output = testOutput
	exit = func(code int) {}
	formatter = new(jsonFormatter)
	constants = make(map[string]interface{})
}
