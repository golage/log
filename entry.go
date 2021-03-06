package log

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
)

const (
	dataValues = "values"
	dataError  = "error"
)

// Entry implements log data
type Entry struct {
	// Raised keeps log raised time
	Raised time.Time

	// Level keeps log level
	Level Level

	// Source keeps log source code with line
	Source string

	// Message keeps log message text
	Message string

	// Data keeps all data
	Data map[string]interface{}
}

// NewEntry returns new entry with defaults
func NewEntry() *Entry {
	var src string
	if pc, file, line, ok := runtime.Caller(2); ok {
		src = fmt.Sprintf("at %v in %v:%d", runtime.FuncForPC(pc).Name(), file, line)
	}
	entry := &Entry{
		Source: src,
		Raised: time.Now(),
		Data:   make(map[string]interface{}),
	}
	for key, value := range constants {
		entry.Data[key] = value
	}
	return entry
}

// Debug logs entry with message in debug level
func (entry Entry) Debug(message string) {
	entry.log(LevelDebug, message)
}

// Info logs entry with message in info level
func (entry Entry) Info(message string) {
	entry.log(LevelInfo, message)
}

// Warning logs entry with message in warning level
func (entry Entry) Warning(message string) {
	entry.log(LevelWarning, message)
}

// Error logs entry with message in error level
func (entry Entry) Error(message string) {
	entry.log(LevelError, message)
}

// Fatal logs entry with message in fatal level so exit with code 1
func (entry Entry) Fatal(message string) {
	entry.log(LevelFatal, message)
	exit(1)
}

// With appends data to entry and returns that
func (entry *Entry) With(data interface{}) *Entry {
	switch value := data.(type) {
	case nil:
		return entry
	case error:
		entry.Data[dataError] = value.Error()
	default:
		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			value = reflect.ValueOf(value).Elem().Interface()
		}
		refType := reflect.TypeOf(value)

		switch refType.Kind() {
		case reflect.Map:
			dic := reflect.ValueOf(value)
			for _, key := range dic.MapKeys() {
				entry.Data[key.String()] = dic.MapIndex(key).Interface()
			}
		case reflect.Struct:
			entry.Data[refType.Name()] = value
		default:
			var values []interface{}
			if entry.Data[dataValues] != nil {
				values = entry.Data[dataValues].([]interface{})
			}
			entry.Data[dataValues] = append(values, value)
		}
	}
	return entry
}

// Value appends key, value to entry and returns that
func (entry *Entry) Value(key string, value interface{}) *Entry {
	entry.Data[key] = value
	return entry
}

func (entry *Entry) log(lvl Level, msg string) {
	entry.Level = lvl
	entry.Message = msg
	if entry.Level >= level {
		if _, err := fmt.Fprintln(output, formatter.Format(*entry)); err != nil {
			log.Printf("can not write on output: %v", err)
		}
	}
}
