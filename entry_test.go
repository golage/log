package log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func createTestEntry() *Entry {
	return &Entry{
		Raised: time.Now(),
		Source: "at test in test.go:11",
		Data:   make(map[string]interface{}),
	}
}

func TestEntry_Debug(t *testing.T) {
	resetTest()
	t.Run("must returns debug message in output", func(t *testing.T) {
		message := "text message"
		createTestEntry().Debug(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelDebug))
	})
}

func TestEntry_Info(t *testing.T) {
	resetTest()
	t.Run("must returns info message in output", func(t *testing.T) {
		message := "text message"
		createTestEntry().Info(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelInfo))
	})
}

func TestEntry_Warning(t *testing.T) {
	resetTest()
	t.Run("must returns warning message in output", func(t *testing.T) {
		message := "text message"
		createTestEntry().Warning(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelWarning))
	})
}

func TestEntry_Error(t *testing.T) {
	resetTest()
	t.Run("must returns error message in output", func(t *testing.T) {
		message := "text message"
		createTestEntry().Error(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelError))
	})
}

func TestEntry_Fatal(t *testing.T) {
	resetTest()
	t.Run("must returns error message in output", func(t *testing.T) {
		message := "text message"
		exit = func(code int) {
			assert.Equal(t, code, 1)
		}
		createTestEntry().Fatal(message)
		text := strings.ToLower(testOutput.String())
		assert.Contains(t, text, fmt.Sprintf("\"message\":\"%v\"", message))
		assert.Contains(t, text, fmt.Sprintf("\"level\":%d", LevelFatal))
	})
}

func TestEntry_Value(t *testing.T) {
	resetTest()
	t.Run("must returns entry with value", func(t *testing.T) {
		entry := createTestEntry().Value("key", "value")
		assert.Equal(t, entry.Data["key"], "value")
	})
}

func TestEntry_With(t *testing.T) {
	resetTest()
	t.Run("must returns entry with value", func(t *testing.T) {
		entry := createTestEntry().With("value1").With("value2").With("value3")
		assert.Equal(t, entry.Data[dataValues], []interface{}{"value1", "value2", "value3"})
	})
	t.Run("must returns entry with pointer value", func(t *testing.T) {
		now := time.Now()
		entry := createTestEntry().With(&now)
		assert.Equal(t, entry.Data["Time"], now)
	})
	t.Run("must returns entry with inline struct", func(t *testing.T) {
		testData := struct {
			Name string
		}{
			Name: "name",
		}
		entry := createTestEntry().With(testData)
		assert.Equal(t, entry.Data[""], testData)
	})
	t.Run("must returns entry with defined struct", func(t *testing.T) {
		type testData struct {
			Name string
		}
		td := testData{Name: "name"}
		entry := createTestEntry().With(td)
		assert.Equal(t, entry.Data["testData"], td)
	})
	t.Run("must returns entry with error", func(t *testing.T) {
		err := fmt.Errorf("test message")
		entry := createTestEntry().With(err)
		assert.Equal(t, entry.Data[dataError], err.Error())
	})
	t.Run("must returns entry with map", func(t *testing.T) {
		dict := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		entry := createTestEntry().With(dict)
		for key, value := range dict {
			assert.Equal(t, entry.Data[key], value)
		}
	})
	t.Run("must returns entry without data affect with nil", func(t *testing.T) {
		entry := createTestEntry().With(nil)
		assert.NotNil(t, entry)
	})
}

func TestNewEntry(t *testing.T) {
	tests := []struct {
		name string
		want *Entry
	}{
		{
			name: "must returns new entry with defaults",
			want: &Entry{
				Raised: time.Now(),
				Source: "at .* in .*[.]go:[\\d]*",
				Data:   make(map[string]interface{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry := NewEntry()
			assert.Equal(t, entry.Data, tt.want.Data)
			assert.Equal(t, entry.Message, tt.want.Message)
			assert.Equal(t, entry.Level, tt.want.Level)
			assert.Equal(t, entry.Raised.Format("2006-01-02 15:04:05"), tt.want.Raised.Format("2006-01-02 15:04:05"))
			assert.Regexp(t, tt.want.Source, entry.Source)
		})
	}
}
