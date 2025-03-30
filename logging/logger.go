package logging

import (
	"log/slog"
	"os"
)

// NewLogger creates a new logger based on the provided level, which must be "debug","info","warn" or "error"
func NewLogger(level string, JSONFormat bool) *Logger {
	var (
		sLevel slog.Level
		opts   *slog.HandlerOptions
		logger *slog.Logger
	)

	if err := sLevel.UnmarshalText([]byte(level)); err != nil {
		slog.Warn("could not unmarshall level, defaulting to info level", "level", level, "error", err)
		sLevel = slog.LevelInfo
	}
	opts = &slog.HandlerOptions{Level: sLevel}

	if JSONFormat {
		logger = slog.New(slog.NewJSONHandler(os.Stderr, opts))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stderr, opts))
	}

	return &Logger{logger}
}

type Logger struct {
	*slog.Logger
}

func (l *Logger) Fatal(msg string, err error) {
	l.Error(msg, "error", err)
	os.Exit(1)
}
