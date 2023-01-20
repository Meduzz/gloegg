package gloegg

import (
	"context"

	"github.com/Meduzz/gloegg/log"
	"github.com/Meduzz/gloegg/toggles"
	"github.com/Meduzz/gloegg/types"
	"github.com/rs/zerolog"
)

var (
	loggers      map[string]zerolog.Logger
	traceHandler func(types.Trace)
)

// Logger will load the named logger or create it.
func Logger(name string) types.Logging {
	if loggers == nil {
		loggers = make(map[string]zerolog.Logger)
	}

	l, exists := loggers[name]

	if !exists {
		l = logger.With().Str("logger", name).Logger()
		loggers[name] = l

		t := toggles.GetToggle(name)

		if t != nil {
			level := t.GetString("info")
			real, err := zerolog.ParseLevel(level)

			if err == nil {
				l.Level(real)
			}
		}
	}

	return log.New(l, traceHandler)
}

// LoggerFromContext will load the named logger or create it. The context are for tracing only.
func LoggerFromContext(name string, ctx context.Context) types.Logging {
	if loggers == nil {
		loggers = make(map[string]zerolog.Logger)
	}

	l, exists := loggers[name]

	if !exists {
		l = logger.With().Str("logger", name).Logger()
		loggers[name] = l

		t := toggles.GetToggle(name)

		if t != nil {
			level := t.GetString(l.GetLevel().String())
			real, err := zerolog.ParseLevel(level)

			if err == nil {
				l.Level(real)
			}
		}
	}

	return log.From(ctx, l, traceHandler)
}

// LoggingToggleListener hook to keep logger levels updated with their toggles
func LoggingToggleListener(t types.Toggle) {
	l, ok := loggers[t.Name()]

	if ok {
		level := t.GetString(l.GetLevel().String())
		real, err := zerolog.ParseLevel(level)

		if err == nil {
			l.Level(real)
		}
	}
}

// SetupLogging allows you to setup the root logger
// any logger created via Logger() or LoggerFromContext()
// will inherit from the root logger.
func SetupLogging(setup func() zerolog.Logger) {
	logger = setup()
}

// SetTraceHandler allows you to set a func to handle
// traces from your app.
func SetTraceHandler(handler func(types.Trace)) {
	traceHandler = handler
}
