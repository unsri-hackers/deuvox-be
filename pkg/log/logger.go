package log

import "github.com/rs/zerolog"

func Error() *zerolog.Event {
	return logger.Error().Stack()
}

func Debug() *zerolog.Event {
	return logger.Debug().Stack()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Fatal() *zerolog.Event {
	return logger.Fatal().Stack()
}

func Panic() *zerolog.Event {
	return logger.Panic().Stack()
}
