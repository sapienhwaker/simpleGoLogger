package logger

import (
	"fmt"
	"os"
)

// Level represents the logging level (e.g., Info, Warning, Error)
type Level int

const (
	LevelInfo  Level = iota // 0
	LevelWarn             // 1
	LevelError            // 2
)

var levelName = map[Level]string{
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
}

// SimpleLogger is a basic logger with different log levels
type SimpleLogger struct {
	level Level
	writer *os.File
}

// NewLogger creates a new SimpleLogger instance
func NewLogger(level Level, logFile string) (*SimpleLogger, error) {
	var writer *os.File
	var err error
	if logFile != "" {
		writer, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
	} else {
		writer = os.Stdout
	}
	return &SimpleLogger{level: level, writer: writer}, nil
}

// SetLevel sets the minimum logging level
func (l *SimpleLogger) SetLevel(level Level) {
	l.level = level
}

// log writes a message to the logger with a specific level
func (l *SimpleLogger) log(level Level, msg string, args ...interface{}) {
	if level >= l.level {
		fmt.Fprintf(l.writer, "[%s] %s\n", levelName[level], fmt.Sprintf(msg, args...))
	}
}

// Info logs an informational message
func (l *SimpleLogger) Info(msg string, args ...interface{}) {
	l.log(LevelInfo, msg, args...)
}

// Warn logs a warning message
func (l *SimpleLogger) Warn(msg string, args ...interface{}) {
	l.log(LevelWarn, msg, args...)
}

// Error logs an error message
func (l *SimpleLogger) Error(msg string, args ...interface{}) {
	l.log(LevelError, msg, args...)
}
