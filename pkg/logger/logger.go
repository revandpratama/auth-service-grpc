package logger

import (
	"time"

	"github.com/rs/zerolog/log"
)

const (
	LEVEL_INFO  = "info"
	LEVEL_WARN  = "warn"
	LEVEL_FATAL = "fatal"
	LEVEL_ERROR = "error"
)

type Logger struct {
	Level     string
	Message   string
	Source    string
	Timestamp time.Time
}

func MakeLog(newLog Logger) {

	switch newLog.Level {
	case LEVEL_INFO:
		log.Info().Msg(newLog.Message)
	case LEVEL_WARN:
		log.Warn().Msg(newLog.Message)
	case LEVEL_ERROR:
		log.Error().Msg(newLog.Message)
	case LEVEL_FATAL:
		log.Fatal().Msg(newLog.Message)
	default:
		log.Info().Msg(newLog.Message)

	}

	// newLog.Timestamp = time.Now()
	// newLog.Source = "auth-service"

	// * Should create kafka here

}
