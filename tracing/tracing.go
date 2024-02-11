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
		checkpoints  []*common.CheckpointDTO
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
		checkpoints:  make([]*common.CheckpointDTO, 0),
	}, nil
}

func NewFromID(id, name string, channel chan *common.Event, logger string, tags ...*common.Tag) (common.Trace, error) {
	traceId, err := tracingid.NewTracingIDFromParent(id, name)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	id, err = tracingid.ToString(traceId)

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
		checkpoints:  make([]*common.CheckpointDTO, 0),
	}, nil
}

func NewFromParent(parent common.Trace, name string, channel chan *common.Event, logger string, tags ...*common.Tag) (common.Trace, error) {
	ctx := context.Background()
	traceId, err := tracingid.NewTracingIDFromParent(parent.ID(), name)

	if err != nil {
		return nil, err
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
		checkpoints:  make([]*common.CheckpointDTO, 0),
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
	trace.Checkpoints = t.checkpoints

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

func (t *tracingImpl) Debug(msg string, tags ...*common.Tag) {
	t.log(common.LevelDebug, msg, tags, nil)
}

func (t *tracingImpl) Error(msg string, err error, tags ...*common.Tag) {
	t.log(common.LevelError, msg, tags, err)
}

func (t *tracingImpl) Info(msg string, tags ...*common.Tag) {
	t.log(common.LevelInfo, msg, tags, nil)
}

func (t *tracingImpl) Warn(msg string, tags ...*common.Tag) {
	t.log(common.LevelWarn, msg, tags, nil)
}

func (l *tracingImpl) log(level, msg string, tags []*common.Tag, err error) {
	log := &common.CheckpointDTO{}
	log.Message = msg
	log.Level = level
	log.Created = time.Now()
	log.Metadata = append(l.tags, tags...)

	if err != nil {
		log.Error = err.Error()
		log.StackTrace = stackMarshaler(err)
	}

	l.checkpoints = append(l.checkpoints, log)
}

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
