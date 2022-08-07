package logger

import (
	"io"
	"os"
	"time"

	"github.com/carpedeez/store/config"
	"github.com/rs/zerolog"
)

func GetLogger(c config.LogConfig) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	var output io.Writer
	if c.Pretty {
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		output = os.Stdout
	}

	return zerolog.New(output).With().Caller().Timestamp().Logger().Level(zerolog.Level(c.Level))
}
