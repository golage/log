package log

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

// Formatter interface of entry to string formatter
type Formatter interface {
	// Format returns entry marshal string from entry
	Format(Entry) string
}

// NewTextFormatter returns new text formatter
func NewTextFormatter() Formatter {
	return new(textFormatter)
}

type textFormatter struct {
}

func (textFormatter) Format(entry Entry) string {
	raised := entry.Raised.Format("2006-01-02 15:04:05.999-07:00")

	level := strings.ToUpper(entry.Level.String())
	switch entry.Level {
	case LevelFatal:
		level = fmt.Sprintf("\033[1;31m%s\033[0m", level)
	case LevelError:
		level = fmt.Sprintf("\033[0;31m%s\033[0m", level)
	case LevelWarning:
		level = fmt.Sprintf("\033[0;33m%s\033[0m", level)
	case LevelInfo:
		level = fmt.Sprintf("\033[0;36m%s\033[0m", level)
	case LevelDebug:
		level = fmt.Sprintf("\033[0;37m%s\033[0m", level)
	}

	raw := fmt.Sprintf("%v | %v | %v \n\t%v", raised, level, strings.TrimSpace(entry.Message), entry.Source)
	if entry.Data != nil && len(entry.Data) > 0 {
		bytes, _ := yaml.Marshal(entry.Data)
		raw = fmt.Sprintf("%s\n\t%s", raw, strings.TrimSpace(strings.ReplaceAll(string(bytes), "\n", "\n\t")))
	}
	return raw
}

// NewJSONFormatter returns new json formatter
func NewJSONFormatter() Formatter {
	return new(jsonFormatter)
}

type jsonFormatter struct {
}

func (jsonFormatter) Format(entry Entry) string {
	bytes, _ := json.Marshal(entry)
	return string(bytes)
}

// NewYAMLFormatter returns new yaml formatter
func NewYAMLFormatter() Formatter {
	return new(yamlFormatter)
}

type yamlFormatter struct {
}

func (yamlFormatter) Format(entry Entry) string {
	bytes, _ := yaml.Marshal(entry)
	return string(bytes)
}
