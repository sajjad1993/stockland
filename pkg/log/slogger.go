package logger

import (
	"context"
	"log/slog"
	"os"
)

type SLogger struct {
	logger *slog.Logger
}

func NewSLogger() Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return &SLogger{
		logger: slog.New(handler),
	}
}

func (l *SLogger) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, args...)
}

func (l *SLogger) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

func (l *SLogger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, args...)
}

func (l *SLogger) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}

func (l *SLogger) Fatal(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
	os.Exit(1)

}

func (l *SLogger) With(ctx context.Context, args ...interface{}) Logger {
	newLogger := l.logger.With(args...)
	return &SLogger{logger: newLogger}
}
