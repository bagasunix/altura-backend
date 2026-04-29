package logging

import (
	"os"
	"strings"
	"time"

	"github.com/phuslu/log"

	"altura-property/internal/infrastructure/config"
)

func Init(cfg *config.Config) {
	var writer log.Writer
	switch cfg.App.Env {
	case "production", "staging":
		writer = &log.IOWriter{Writer: os.Stdout}
	default:
		writer = &log.ConsoleWriter{
			ColorOutput: true,
			QuoteString: true,
			Writer:      os.Stdout,
		}
	}

	log.DefaultLogger = log.Logger{
		Level:      parseLevel(cfg.Logger.Level),
		Caller:     1,
		TimeField:  "ts",
		TimeFormat: time.RFC3339Nano,
		Writer:     writer,
		Context: log.NewContext(nil).
			Str("service", cfg.App.Name).
			Str("env", cfg.App.Env).
			Value(),
	}
}

func parseLevel(s string) log.Level {
	switch strings.ToLower(s) {
	case "debug":
		return log.DebugLevel
	case "warn", "warning":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	default:
		return log.InfoLevel
	}
}
