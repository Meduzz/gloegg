package common

import "time"

type (
	Event struct {
		Kind     string    `json:"kind"`     // LOG|TRACE
		Logger   string    `json:"logger"`   // logger
		Metadata []*Tag    `json:"metadata"` // metadata
		Created  time.Time `json:"created"`
		Log      *LogDTO   `json:"log,omitempty"`
		Trace    *TraceDTO `json:"trace,omitempty"`
	}

	LogDTO struct {
		Level      string   `json:"level"`
		Message    string   `json:"message"`
		Error      string   `json:"error,omitempty"` // error from error logs and traces
		StackTrace []*Stack `json:"stack,omitempty"` // Stack trace from error log
	}

	CheckpointDTO struct {
		Level      string    `json:"level"`
		Message    string    `json:"message"`
		Error      string    `json:"error,omitempty"` // error from error logs and traces
		StackTrace []*Stack  `json:"stack,omitempty"` // Stack trace from error log
		Created    time.Time `json:"created"`
		Metadata   []*Tag    `json:"metadata"`
	}

	TraceDTO struct {
		Name        string           `json:"name"`
		ID          string           `json:"id"`
		Start       time.Time        `json:"start"`
		End         time.Time        `json:"end"`
		Checkpoints []*CheckpointDTO `json:"checkpoints"`
		Error       string           `json:"error,omitempty"` // error from error logs and traces
		StackTrace  []*Stack         `json:"stack,omitempty"` // Stack trace from error log
	}

	Sink interface {
		Handle(*Event)
	}
)
