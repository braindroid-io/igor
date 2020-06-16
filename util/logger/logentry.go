package logger

import (
	"encoding/json"
)

type LogEntry struct {
	Message string `json:"message"`
	Level string `json:"level"`
}

func (entry *LogEntry) Serialize() ([]byte, error) {
	b, err := json.Marshal(entry)
	return b, err
}
