package common

import "time"

type (
	Event struct {
		Kind     string    `json:"kind"`     // LOG|TRACE
		Logger   string    `json:"logger"`   // logger
		Metadata []*Tag    `json:"metadata"` // metadata
		Created  time.Time `json:"created"`
		Log      *LogDTO   `json:"log"`
		Trace    *TraceDTO `json:"trace"`
	}

	LogDTO struct {
		Level      string   `json:"level"`
		Message    string   `json:"message"`
		Error      error    `json:"error"` // error from error logs and traces
		StackTrace []*Stack `json:"stack"` // Stack trace from error log
	}

	TraceDTO struct {
		Name       string    `json:"name"`
		ID         string    `json:"id"`
		Start      time.Time `json:"start"`
		End        time.Time `json:"end"`
		Error      error     `json:"error,omitempty"` // error from error logs and traces
		StackTrace []*Stack  `json:"stack,omitempty"` // Stack trace from error log
	}

	Sink interface {
		Handle(*Event)
	}
)
