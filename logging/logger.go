package logging

import (
	"log/slog"
	"os"
)

func NewLogger() Logger {
	level := slog.LevelDebug
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
	return Logger{s: logger}
}

type Logger struct {
	s *slog.Logger
}

func (l Logger) Fatal(msg string) {
	l.s.Error(msg)
	os.Exit(1)
}
