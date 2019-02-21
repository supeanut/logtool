package lager

import (
	"encoding/json"
	"github.com/supeanut/logtool/lager/openlogging"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func FormatLogLevel(x LogLevel) string {
	var level string
	switch x {
	case DEBUG:
		level = "DEBUG"
	case INFO:
		level = "INFO"
	case WARN:
		level = "WARN"
	case ERROR:
		level = "ERROR"
	case FATAL:
		level = "FATAL"
	}
	return level
}

func (x LogLevel) MarshalJSON() ([]byte, error) {
	var level = FormatLogLevel(x)
	return json.Marshal(level)
}

type LogFormat struct {
	LogLevel LogLevel `json:"level"`
	Timestamp string `json:"timestamp"`
	File string `json:"file"`
	Message string `json:"msg"`
	Data openlogging.Tags `json:"data,omitempty"`
}

func (log LogFormat) ToJSON() ([]byte, error) {
	return json.Marshal(log)
}
