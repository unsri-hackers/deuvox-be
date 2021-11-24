package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func New() {
	env := os.Getenv("GO_ENV")
	var w io.Writer
	// TODO: add writer for production
	if env == "dev" {
		w = zerolog.ConsoleWriter{Out: os.Stdout}
	} else {
		w = os.Stdout
	}
	l := zerolog.New(w).With().Timestamp().Caller().Str("env", env).Logger()
	logger = l
}
