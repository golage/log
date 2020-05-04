package log

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
	"time"
)

func TestNewJSONFormatter(t *testing.T) {
	tests := []struct {
		name string
		want Formatter
	}{
		{
			name: "must returns json formatter",
			want: new(jsonFormatter),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewJSONFormatter(), tt.want)
		})
	}
}

func TestNewTextFormatter(t *testing.T) {
	tests := []struct {
		name string
		want Formatter
	}{
		{
			name: "must returns text formatter",
			want: new(textFormatter),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewTextFormatter(), tt.want)
		})
	}
}

func TestNewYAMLFormatter(t *testing.T) {
	tests := []struct {
		name string
		want Formatter
	}{
		{
			name: "must returns yaml formatter",
			want: new(yamlFormatter),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewYAMLFormatter(), tt.want)
		})
	}
}

func Test_textFormatter_Format(t *testing.T) {
	formatter := new(textFormatter)
	type Test struct {
		name string
		arg  Entry
	}
	var tests []Test
	for _, lvl := range testLevels {
		tests = append(tests, Test{
			name: fmt.Sprintf("must returns marshal string with %s entry", lvl),
			arg: Entry{
				Raised:  time.Now(),
				Level:   lvl,
				Source:  "at test in test.go:10",
				Message: "text message",
				Data: map[string]interface{}{
					"error": "can not do job",
					"key":   "value",
				},
			},
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text := strings.ToLower(formatter.Format(tt.arg))
			assert.Contains(t, text, strings.ToLower(tt.arg.Level.String()))
			assert.Contains(t, text, strings.ToLower(tt.arg.Source))
			assert.Contains(t, text, strings.ToLower(tt.arg.Message))
			assert.Contains(t, text, strings.ToLower(tt.arg.Raised.Format("2006-01-02 15:04:05")))
			bytes, _ := yaml.Marshal(tt.arg.Data)
			for _, line := range strings.Split(string(bytes), "\n") {
				assert.Contains(t, text, strings.ToLower(strings.TrimSpace(line)))
			}
		})
	}
}

func Test_jsonFormatter_Format(t *testing.T) {
	formatter := new(jsonFormatter)
	tests := []struct {
		name string
		arg  Entry
	}{
		{
			name: "must returns json marshal",
			arg: Entry{
				Raised:  time.Now(),
				Level:   LevelWarning,
				Source:  "at test in test.go:10",
				Message: "text message",
				Data: map[string]interface{}{
					"error": "can not do job",
					"key":   "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := json.Marshal(tt.arg)
			assert.Equal(t, formatter.Format(tt.arg), string(bytes))
		})
	}
}

func Test_yamlFormatter_Format(t *testing.T) {
	formatter := new(yamlFormatter)
	tests := []struct {
		name string
		arg  Entry
	}{
		{
			name: "must returns yaml marshal",
			arg: Entry{
				Raised:  time.Now(),
				Level:   LevelWarning,
				Source:  "at test in test.go:10",
				Message: "text message",
				Data: map[string]interface{}{
					"error": "can not do job",
					"key":   "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := yaml.Marshal(tt.arg)
			assert.Equal(t, formatter.Format(tt.arg), string(bytes))
		})
	}
}
