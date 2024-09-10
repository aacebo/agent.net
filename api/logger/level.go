package logger

import (
	"log/slog"
	"strings"

	"github.com/aacebo/agent.net/api/utils"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
)

func GetEnvLevel() (Level, bool) {
	level := Level(strings.ToLower(utils.GetEnv("LOG_LEVEL", string(Debug))))
	return level, level.Valid()
}

func (self Level) Valid() bool {
	switch self {
	case Debug, Info, Warn, Error:
		return true
	}

	return false
}

func (self Level) SLog() slog.Level {
	switch self {
	case Debug:
		return slog.LevelDebug
	case Info:
		return slog.LevelInfo
	case Warn:
		return slog.LevelWarn
	case Error:
		return slog.LevelError
	}

	panic("invalid log level")
}
