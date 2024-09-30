package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/aacebo/agent.net/core/utils"
)

func New(name string) *slog.Logger {
	prefix := os.Getenv("LOG_PREFIX")

	if prefix != "" {
		prefix = prefix + "/"
	}

	lvl := slog.LevelVar{}

	if level := Level(strings.ToLower(utils.GetEnv("LOG_LEVEL", string(Debug)))); level.Valid() {
		lvl.Set(level.SLog())
	}

	return slog.New(NewColorTextHandler(&slog.HandlerOptions{
		Level:     &lvl,
		AddSource: lvl.Level() == slog.LevelDebug,
	})).With("name", prefix+name)
}
