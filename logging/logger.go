package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/tracing"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/go-stack/stack"
)

type (
	loggingFacade struct {
		eventChannel chan *common.Event
		metadata     []*common.Tag
		name         string
	}
)

func NewLogger(name string, channel chan *common.Event, systemMeta []*common.Tag) common.Logger {
	return &loggingFacade{channel, systemMeta, name}
}

func (l *loggingFacade) Info(msg string, tags ...*common.Tag) {
	l.log(common.LevelInfo, msg, tags, nil)
}

func (l *loggingFacade) Debug(msg string, tags ...*common.Tag) {
	l.log(common.LevelDebug, msg, tags, nil)
}

func (l *loggingFacade) Warn(msg string, tags ...*common.Tag) {
	l.log(common.LevelWarn, msg, tags, nil)
}

func (l *loggingFacade) Error(msg string, err error, tags ...*common.Tag) {
	l.log(common.LevelError, msg, tags, err)
}

func (l *loggingFacade) Trace(name string, tags ...*common.Tag) common.Trace {
	return tracing.New(name, l.eventChannel, l.name, append(l.metadata, tags...)...)
}

func (l *loggingFacade) TraceContext(name string, parent context.Context, tags ...*common.Tag) (common.Trace, error) {
	return tracing.NewFromContext(name, parent, l.eventChannel, l.name, append(l.metadata, tags...)...)
}

func (l *loggingFacade) log(level, msg string, tags []*common.Tag, err error) {
	log := &common.LogDTO{}
	log.Message = msg
	log.Level = level

	event := &common.Event{}
	event.Kind = "LOG"
	event.Metadata = slice.Concat(l.metadata, tags)
	event.Created = time.Now()
	event.Logger = l.name

	if err != nil {
		log.Error = err.Error()
		log.StackTrace = stackMarshaler(err)
	}

	event.Log = log

	l.eventChannel <- event
}

// stackMarshaler excels at marshaling stack traces from logs
func stackMarshaler(e error) []*common.Stack {
	arr := make([]*common.Stack, 0)

	it := stack.Trace().TrimRuntime()[3:]

	for _, trace := range it {
		frame := trace.Frame()
		arr = append(arr, &common.Stack{
			File: fmt.Sprintf("%+s", trace),
			Line: frame.Line,
			Func: frame.Function,
		})
	}

	return arr
}
