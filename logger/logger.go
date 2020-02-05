package logger

import "fmt"

// Logger logger
type Logger struct {
	Level int32 `json:"level,omitemtpy"` // 等级
}

// NewLogger new logger
func NewLogger() *Logger {
	return &Logger{}
}

// Info log info l.Level 1
func (l *Logger) Info() {
	if l.Level <= 1 {
		fmt.Println(l.Level, "Info")
	}
}

// Debug log debug l.Level 2
func (l *Logger) Debug() {
	if l.Level <= 2 {
		fmt.Println(l.Level, "Debug")
	}
}

// Warn log warn l.Level 3
func (l *Logger) Warn() {
	if l.Level <= 3 {
		fmt.Println(l.Level, "Warn")
	}
}

// Error log error l.Level 4
func (l *Logger) Error() {
	if l.Level <= 4 {
		fmt.Println(l.Level, "Error")
	}
}

// Fatal log fatal l.Level 5
func (l *Logger) Fatal() {
	if l.Level <= 5 {
		fmt.Println(l.Level, "Fatal")
	}
}
