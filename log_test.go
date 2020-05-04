package log

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDebug(t *testing.T) {
	resetTest()
	t.Run("must returns debug message in output", func(t *testing.T) {
		message := "text message"
		Debug(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelDebug))
	})
}

func TestInfo(t *testing.T) {
	resetTest()
	t.Run("must returns info message in output", func(t *testing.T) {
		message := "text message"
		Info(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelInfo))
	})
}

func TestWarning(t *testing.T) {
	resetTest()
	t.Run("must returns warning message in output", func(t *testing.T) {
		message := "text message"
		Warning(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelWarning))
	})
}

func TestError(t *testing.T) {
	resetTest()
	t.Run("must returns error message in output", func(t *testing.T) {
		message := "text message"
		Error(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelError))
	})
}

func TestFatal(t *testing.T) {
	resetTest()
	t.Run("must returns fatal message in output", func(t *testing.T) {
		message := "fatal text message"
		exit = func(code int) {
			assert.Equal(t, code, 1)
		}
		Fatal(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelFatal))
	})
}

func TestSetConstant(t *testing.T) {
	resetTest()
	data := "\"Data\":{\"const_key\":\"const_value\"}"
	SetConstant("const_key", "const_value")
	for _, lvl := range testLevels {
		t.Run(fmt.Sprintf("must returns constant values in %v", strings.ToLower(lvl.String())), func(t *testing.T) {
			switch lvl {
			case LevelDebug:
				Debug("")
			case LevelInfo:
				Info("")
			case LevelWarning:
				Warning("")
			case LevelError:
				Error("")
			case LevelFatal:
				Fatal("")
			}
			assert.Contains(t, testOutput.String(), data)
			testOutput.Reset()
		})
	}
}

func TestSetLevel(t *testing.T) {
	resetTest()
	for _, lvl := range testLevels {
		SetLevel(lvl)
		t.Run(fmt.Sprintf("must allowed to logs equal or greater than %v", strings.ToLower(lvl.String())), func(t *testing.T) {
			for _, l := range testLevels {
				switch l {
				case LevelDebug:
					Debug("")
				case LevelInfo:
					Info("")
				case LevelWarning:
					Warning("")
				case LevelError:
					Error("")
				case LevelFatal:
					Fatal("")
				}
				if lvl <= l {
					assert.NotEmpty(t, testOutput.String())
				} else {
					assert.Empty(t, testOutput.String())
				}
				testOutput.Reset()
			}
		})
	}
}

func TestSetFormatter(t *testing.T) {
	resetTest()
	t.Run("must sets json-formatter in formatter var", func(t *testing.T) {
		f := new(jsonFormatter)
		SetFormatter(f)
		assert.Equal(t, formatter, f)
	})
	t.Run("must sets yaml-formatter in formatter var", func(t *testing.T) {
		f := new(yamlFormatter)
		SetFormatter(f)
		assert.Equal(t, formatter, f)
	})
	t.Run("must sets text-formatter in formatter var", func(t *testing.T) {
		f := new(textFormatter)
		SetFormatter(f)
		assert.Equal(t, formatter, f)
	})
}

func TestSetOutput(t *testing.T) {
	resetTest()
	t.Run("must sets buffer in output var", func(t *testing.T) {
		buf := new(bytes.Buffer)
		SetOutput(buf)
		assert.Equal(t, output, buf)
	})
}

func TestValue(t *testing.T) {
	resetTest()
	t.Run("must returns entry with value", func(t *testing.T) {
		entry := Value("key", "value")
		assert.Equal(t, entry.Data["key"], "value")
	})
}

func TestWith(t *testing.T) {
	resetTest()
	t.Run("must returns entry with value", func(t *testing.T) {
		entry := With("value1").With("value2").With("value3")
		assert.Equal(t, entry.Data[dataValues], []interface{}{"value1", "value2", "value3"})
	})
	t.Run("must returns entry with inline struct", func(t *testing.T) {
		testData := struct {
			Name string
		}{
			Name: "name",
		}
		entry := With(testData)
		assert.Equal(t, entry.Data[""], testData)
	})
	t.Run("must returns entry with defined struct", func(t *testing.T) {
		type testData struct {
			Name string
		}
		td := testData{Name: "name"}
		entry := With(td)
		assert.Equal(t, entry.Data["testData"], td)
	})
	t.Run("must returns entry with error", func(t *testing.T) {
		err := fmt.Errorf("test message")
		entry := With(err)
		assert.Equal(t, entry.Data[dataError], err.Error())
	})
	t.Run("must returns entry with map", func(t *testing.T) {
		dict := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		entry := With(dict)
		for key, value := range dict {
			assert.Equal(t, entry.Data[key], value)
		}
	})
	t.Run("must returns entry without data affect with nil", func(t *testing.T) {
		entry := With(nil)
		assert.NotNil(t, entry)
	})
}
