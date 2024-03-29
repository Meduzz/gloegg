package common

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/Meduzz/helper/fp/slice"
)

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
		Parent      string           `json:"parent,omitempty"` // parent trace if any
		Start       time.Time        `json:"start"`
		End         time.Time        `json:"end"`
		Checkpoints []*CheckpointDTO `json:"checkpoints"`
		Error       string           `json:"error,omitempty"` // error from error logs and traces
		StackTrace  []*Stack         `json:"stack,omitempty"` // Stack trace from error log
	}

	Sink interface {
		Handle(*Event)
		Name() string // identify this sink among others (used in removal)
	}

	LogFormat func(*Event) string
)

func DefaultFormat(event *Event) string {
	created := event.Created.Format(time.DateTime)
	logger := event.Logger

	buf := bytes.NewBufferString("")

	if event.Kind == "LOG" {
		// ts [logger] level - message (metadata)
		level := event.Log.Level

		fmt.Fprintf(buf, "%s [%s] %s - %s [%s]\n", created, logger, level, event.Log.Message, dumpMetadata(event.Metadata))

		if event.Log.Error != "" {
			// <error message>
			fmt.Fprintf(buf, "Error:\n\t%s\n", event.Log.Error)
			fmt.Fprintln(buf, "Stacktrace:")

			// func - file:line
			slice.ForEach(event.Log.StackTrace, func(item *Stack) {
				// TODO filter external toggle, that prefix check the item.Func to get rid of noise?
				fmt.Fprintf(buf, "\t%s - %s:%d\n", item.Func, item.File, item.Line)
			})
		}
	} else if event.Kind == "TRACE" {
		// ts [logger] trace: message (duration?) (metadata)
		start := event.Trace.Start
		end := event.Trace.End

		fmt.Printf("%s [%s] trace: %s (%s) [%s]\n", created, logger, event.Trace.Name, end.Sub(start).String(), dumpMetadata(event.Metadata))

		slice.ForEach(event.Trace.Checkpoints, func(it *CheckpointDTO) {
			// ts [logger] checkpoint - message (duration) metadata
			fmt.Fprintf(buf, "%s [%s/%s] checkpoint (%s) - %s [%s]\n", it.Created.Format(time.DateTime), logger, event.Trace.Name, it.Level, it.Message, dumpMetadata(it.Metadata))

			if it.Error != "" {
				// <error message>
				fmt.Fprintf(buf, "Error:\n\t%s\n", it.Error)
				fmt.Fprintln(buf, "Stacktrace:")

				// func - file:line
				slice.ForEach(it.StackTrace, func(item *Stack) {
					// TODO filter external toggle, that prefix check the item.Func to get rid of noise?
					fmt.Fprintf(buf, "\t%s - %s:%d\n", item.Func, item.File, item.Line)
				})
			}

		})
	}

	return buf.String()
}

func dumpMetadata(tags []*Tag) string {
	converted := slice.Map(tags, func(tag *Tag) string {
		return fmt.Sprintf("%s=%v", tag.Key, tag.Value)
	})

	return strings.Join(converted, ", ")
}
