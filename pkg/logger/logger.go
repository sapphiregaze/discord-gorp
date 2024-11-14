package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

var logger *slog.Logger

type logFunc func(msg string, args ...any)

func makeLogFunc(level slog.Level) logFunc {
	return func(msg string, args ...any) {
		logger.Log(nil, level, msg, args...)
	}
}

var (
	Info  = makeLogFunc(slog.LevelInfo)
	Error = makeLogFunc(slog.LevelError)
	Warn  = makeLogFunc(slog.LevelWarn)
	Debug = makeLogFunc(slog.LevelDebug)
)

func init() {
	w := os.Stderr
	logger = slog.New(tint.NewHandler(w, nil))
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
}
