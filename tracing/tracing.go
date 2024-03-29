package tracing

import (
	"fmt"
	"time"

	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/helper/hashing"
	"github.com/go-stack/stack"
)

type (
	tracingImpl struct {
		id           string
		parent       string
		name         string
		start        time.Time
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
)

const (
	GloeggTraceKey string = "GLOEGG_TRACE_ID"
)

var (
	ErrUnreadableTraceID = fmt.Errorf("could not read the provided traceID")
)

func New(name string, channel chan *common.Event, logger string, tags ...*common.Tag) common.Trace {
	return &tracingImpl{
		id:           hashing.Token(),
		parent:       "",
		logger:       logger,
		name:         name,
		tags:         tags,
		start:        time.Now(),
		eventChannel: channel,
		done:         false,
		checkpoints:  make([]*common.CheckpointDTO, 0),
	}
}

func NewFromID(parent, name string, channel chan *common.Event, logger string, tags ...*common.Tag) (common.Trace, error) {
	return &tracingImpl{
		id:           hashing.Token(),
		parent:       parent,
		logger:       logger,
		name:         name,
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

func (t *tracingImpl) Parent() string {
	return t.parent
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
	event.Kind = common.KindTrace
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
