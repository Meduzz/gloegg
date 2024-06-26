package common

type (
	Logger interface {
		// Info - log a msg at info level, providing additional context via tags
		Info(msg string, tags ...*Tag)

		// Debug - log a msg at debug level, providing additional context via tags
		Debug(msg string, tags ...*Tag)

		// Warn - log a msg at warn level, providing additional context via tags
		Warn(msg string, tags ...*Tag)

		// Error - log a msg at error level, providing additional context via tags
		Error(msg string, err error, tags ...*Tag)

		// Trace - create a new named trace, providing additional context via tags
		Trace(name string, tags ...*Tag) Trace

		// TraceParentID - create a new named trace from a parent trace id.
		TraceFromParentID(parent, name string, tags ...*Tag) (Trace, error)
	}

	Trace interface {
		// Fetch the id of this trace
		ID() string

		Parent() string

		// Complete this trace with an optional error
		Done(error)

		// Add additional metadata to this trace
		AddMetadata(...*Tag)

		// Info - log a msg at info level, providing additional context via tags
		Info(msg string, tags ...*Tag)

		// Debug - log a msg at debug level, providing additional context via tags
		Debug(msg string, tags ...*Tag)

		// Warn - log a msg at warn level, providing additional context via tags
		Warn(msg string, tags ...*Tag)

		// Error - log a msg at error level, providing additional context via tags
		Error(msg string, err error, tags ...*Tag)
	}

	Tag struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}

	Stack struct {
		File string `json:"file"`
		Line int    `json:"line"`
		Func string `json:"function"`
	}
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	KindLog    = "LOG"
	KindTrace  = "TRACE"
)
