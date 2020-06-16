package logger

import (
	"log"
	"fmt"
)

type Logger struct {}

func (logger *Logger) Log(entry *LogEntry) {
	b, err := entry.Serialize()
	if err != nil {
		log.Printf("error: %s", err.Error())
	}

	log.Printf(string(b))
}

func (logger *Logger) Info(message string, args ...interface{}) {
	entry := &LogEntry{
		Message: fmt.Sprintf(message, args...),
		Level: "Info",
	}

	logger.Log(entry)
}

func (logger *Logger) Debug(message string, args ...interface{}) {
	entry := &LogEntry{
		Message: fmt.Sprintf(message, args...),
		Level: "Debug",
	}

	logger.Log(entry)
}

func (logger *Logger) Error(message string, args ...interface{}) {
	entry := &LogEntry{
		Message: fmt.Sprintf(message, args...),
		Level: "Error",
	}

	logger.Log(entry)
}

func (logger *Logger) Fatal(message string, args ...interface{}) {
	entry := &LogEntry{
		Message: fmt.Sprintf(message, args...),
		Level: "Fatal",
	}

	logger.Log(entry)
}
