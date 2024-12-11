package logging

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func Log(level zerolog.Level, callerName string, msg string) {
	callerName = strings.ToUpper(callerName)
	Logger.WithLevel(level).Msgf(" | [%s] | %s", callerName, msg)
}

func Logf(level zerolog.Level, callerName string, format string, v ...any) {
	callerName = strings.ToUpper(callerName)
	Logger.WithLevel(level).Msgf(" | ["+callerName+"] | "+format, v...)
}

func Init() {
	// Zerolog configuration
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i any) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
	}
	Logger = zerolog.New(output).With().Timestamp().Logger()
}
