package logger

import (
	"os"
	"project/go-fiber/config"

	"github.com/rs/zerolog"
)

func NewZeroSlogLogger(config *config.LoggerConfig) *zerolog.Logger {
	var logger zerolog.Logger
	zerolog.SetGlobalLevel(zerolog.Level(config.Level))
	if config.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}
	return &logger
}
