package logging

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func Log(level zerolog.Level, callerName string, msg string) {
	callerName = strings.ToUpper(callerName)
	Logger.WithLevel(level).Msgf("| [%s] | %s", callerName, msg)
}

func Logf(level zerolog.Level, callerName string, format string, v ...any) {
	callerName = strings.ToUpper(callerName)
	Logger.WithLevel(level).Msgf("| ["+callerName+"] | "+format, v...)
}

func Init() {
	// Zerolog configuration
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	Logger = zerolog.New(output).With().Timestamp().Logger()
}
