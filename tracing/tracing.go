package tracing

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/tracing/tracingid"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/go-stack/stack"
)

type (
	tracingImpl struct {
		id           string
		name         string
		start        time.Time
		ctx          context.Context
		tags         []*common.Tag
		eventChannel chan *common.Event
		done         bool
		logger       string
	}

	TracingError struct {
		Error string          `json:"error"`
		Stack []*common.Stack `json:"stack"`
	}

	contextKey struct {
		key string
	}
)

const (
	GloeggTraceKey string = "GLOEGG_TRACE_ID"
)

var (
	ErrUnreadableTraceID = fmt.Errorf("could not read the provided traceID")
)

func New(name string, channel chan *common.Event, logger string, tags ...*common.Tag) common.Trace {
	it, _ := NewFromContext(name, nil, channel, logger, tags...)

	return it
}

func NewFromContext(name string, ctx context.Context, channel chan *common.Event, logger string, tags ...*common.Tag) (common.Trace, error) {
	var traceId *tracingid.TraceId

	if ctx == nil {
		ctx = context.Background()
		traceId = tracingid.NewTracingID(name)
	} else {
		p, ok := ctx.Value(&contextKey{GloeggTraceKey}).(string)

		if ok {
			newId, err := tracingid.NewTracingIDFromParent(p, name)

			if err != nil {
				return nil, errors.Join(ErrUnreadableTraceID, err)
			}

			traceId = newId
		} else {
			traceId = tracingid.NewTracingID(name)
		}
	}

	id, err := tracingid.ToString(traceId)

	if err != nil {
		return nil, err
	}

	return &tracingImpl{
		id:           id,
		logger:       logger,
		name:         name,
		ctx:          context.WithValue(ctx, &contextKey{GloeggTraceKey}, traceId),
		tags:         tags,
		start:        time.Now(),
		eventChannel: channel,
		done:         false,
	}, nil
}

func (t *tracingImpl) ID() string {
	return t.id
}

func (t *tracingImpl) Done(err error) {
	if t.done {
		return
	}

	trace := &common.TraceDTO{}
	trace.End = time.Now()
	trace.Name = t.name
	trace.Start = t.start
	trace.ID = t.id

	event := &common.Event{}
	event.Created = time.Now()
	event.Kind = "TRACE"
	event.Metadata = t.tags
	event.Logger = t.logger

	it := stack.Trace().TrimRuntime()[1:]

	callstack := slice.Map(it, func(call stack.Call) *common.Stack {
		frame := call.Frame()

		return &common.Stack{
			File: fmt.Sprintf("%+s", call),
			Line: frame.Line,
			Func: frame.Function,
		}
	})

	if err != nil {
		trace.Error = err.Error()
		trace.StackTrace = callstack
	}

	event.Trace = trace

	t.eventChannel <- event
	t.done = true
}

func (t *tracingImpl) Context() context.Context {
	return t.ctx
}

func (t *tracingImpl) AddMetadata(tags ...*common.Tag) {
	t.tags = append(t.tags, tags...)
}
