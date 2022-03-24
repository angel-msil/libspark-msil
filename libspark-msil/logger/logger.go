package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"libspark-msil/constants"
	"fmt"
)

type Level string

func (l Level) zeroLogLevel() zerolog.Level {
	switch l {
	case constants.TraceLevel:
		return zerolog.TraceLevel
	case constants.DebugLevel:
		return zerolog.DebugLevel
	case constants.InfoLevel:
		return zerolog.InfoLevel
	case constants.WarnLevel:
		return zerolog.WarnLevel
	case constants.ErrorLevel:
		return zerolog.ErrorLevel
	case constants.FatalLevel:
		return zerolog.FatalLevel
	case constants.PanicLevel:
		return zerolog.PanicLevel
	default:
		return zerolog.DebugLevel
	}
}

// StartLogger is used to configure the logger and make it ready for use
func StartLogger(level Level) {
fmt.Println("start logger func call")
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(level.zeroLogLevel())
	fmt.Println("start logger func call 3")
	log.Logger = log.With().Caller().Logger()
	fmt.Println("start logger func call 4")
	log.Info().Str(constants.LogLevelKey, zerolog.GlobalLevel().String()).Msg("started logger")
	fmt.Println("start logger func call started")
}
