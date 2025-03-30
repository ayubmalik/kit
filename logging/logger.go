package logging

import (
	"log/slog"
	"os"
)

// NewLogger creates a new logger based on the provided level, which must be "debug","info","warn" or "error"
func NewLogger(level string) Logger {
	var l slog.Level
	if err := l.UnmarshalText([]byte(level)); err != nil {
		slog.Warn("could not unmarshall level, defaulting to 'info'", "level", level, "error", err)
		l = slog.LevelInfo
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: l}))
	return Logger{s: logger}
}

type Logger struct {
	s *slog.Logger
}

func (l Logger) Fatal(msg string, err error) {
	l.s.Error(msg, "error", err)
	os.Exit(1)
}
