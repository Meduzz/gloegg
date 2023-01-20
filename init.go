package gloegg

import (
	"fmt"

	"github.com/Meduzz/gloegg/toggles"
	"github.com/Meduzz/gloegg/types"
	"github.com/go-stack/stack"
	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
)

func init() {
	SetupLogging(SetupDefaultLogging)
	toggles.SetToggleUpdatedHandler(LoggingToggleListener)
}

// LoggingDefaults will set TimeFieldFormat to TimeFormatUnixMs
func LoggingDefaults() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.ErrorStackMarshaler = StackMarshaler
}

// SetupDefaultLogging sets up a default logger by calling
// LogginDefaults and creating a new info logger with console writer
func SetupDefaultLogging() zerolog.Logger {
	LoggingDefaults()

	console := zerolog.NewConsoleWriter()

	logger := zerolog.New(console).With().Stack().Timestamp().Logger()
	return logger.Level(zerolog.InfoLevel)
}

// MultiToggleUpdatedListener combines several listeners into one
func MultiToggleUpdatedListener(handlers ...func(types.Toggle)) func(types.Toggle) {
	return func(t types.Toggle) {
		for _, it := range handlers {
			it(t)
		}
	}
}

// StackMarshaler excels at marshaling stack traces from logs
func StackMarshaler(e error) interface{} {
	arr := make([]*types.Stack, 0)

	it := stack.Trace().TrimRuntime()[3:]

	for _, trace := range it {
		frame := trace.Frame()
		arr = append(arr, &types.Stack{
			File: fmt.Sprintf("%+s", trace),
			Line: frame.Line,
			Func: frame.Function,
		})
	}

	return arr
}
