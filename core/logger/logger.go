package logger

import (
	"log/slog"
)

func New(name string) *slog.Logger {
	lvl := slog.LevelVar{}

	if v, ok := GetEnvLevel(); ok {
		lvl.Set(v.SLog())
	}

	return slog.New(NewColorTextHandler(&slog.HandlerOptions{
		Level:     &lvl,
		AddSource: lvl.Level() == slog.LevelDebug,
	})).With("name", name)
}
