package log

import (
	"io"
	"os"
)

var (
	output    io.Writer = os.Stdout
	level               = LevelInfo
	exit                = os.Exit
	formatter           = NewTextFormatter()
	constants           = make(map[string]interface{})
)

// SetOutput sets logging output
func SetOutput(w io.Writer) {
	output = w
}

// SetLevel sets logging minimum level
func SetLevel(lvl Level) {
	level = lvl
}

// SetFormatter sets logging formatter
func SetFormatter(f Formatter) {
	formatter = f
}

// SetConstant sets logging constants data
func SetConstant(key string, value interface{}) {
	constants[key] = value
}

// Debug creates entry with message and logs in debug level
func Debug(message string) {
	NewEntry().Debug(message)
}

// Info creates entry with message and logs in info level
func Info(message string) {
	NewEntry().Info(message)
}

// Warning creates entry with message and logs in warning level
func Warning(message string) {
	NewEntry().Warning(message)
}

// Error creates entry with message and logs in error level
func Error(message string) {
	NewEntry().Error(message)
}

// Fatal creates entry with message and logs in fatal level so exit with code 1
func Fatal(message string) {
	NewEntry().Fatal(message)
}

// With creates entry with data and returns that
func With(data interface{}) *Entry {
	return NewEntry().With(data)
}

// Value creates entry with key, value and returns that
func Value(key string, value interface{}) *Entry {
	return NewEntry().Value(key, value)
}
