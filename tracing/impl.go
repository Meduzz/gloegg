package tracing

import (
	"context"
	"fmt"
	"time"

	"github.com/Meduzz/gloegg/types"
	"github.com/Meduzz/helper/hashing"
	"github.com/go-stack/stack"
)

type (
	tracingImpl struct {
		Id        string `json:"id"`
		Parent    string `json:"parent,omitempty"`
		NameField string `json:"name"`
		StartTS   int64  `json:"start"`
		EndTS     int64  `json:"end"`
		ctx       context.Context
		Tags      map[string]interface{} `json:"tags"`
		Err       *TracingError          `json:"error"`
		handler   func(types.Trace)
	}

	TracingError struct {
		Error string         `json:"error"`
		Stack []*types.Stack `json:"stack"`
	}

	contextKey struct {
		key string
	}
)

const (
	GloeggTraceKey    string = "GLOEGG_TRACE"
	GloeggTraceHeader string = "x-gloegg-trace"
)

func New(name string, ctx context.Context, handler func(types.Trace), tags ...*types.Tag) types.Trace {
	id := hashing.Token()
	parent := ""

	if ctx == nil {
		ctx = context.Background()
	} else {
		p, ok := ctx.Value(&contextKey{GloeggTraceKey}).(string)

		if ok {
			parent = p
		}
	}

	return &tracingImpl{
		Id:        id,
		Parent:    parent,
		NameField: name,
		ctx:       context.WithValue(ctx, &contextKey{GloeggTraceKey}, id),
		Tags:      types.AsMap(tags...),
		handler:   handler,
	}
}

func (t *tracingImpl) Start() (context.Context, func(error)) {
	t.StartTS = time.Now().UnixMilli()

	return t.ctx, t.stopper
}

func (t *tracingImpl) ID() string {
	return t.Id
}

func (t *tracingImpl) Name() string {
	return t.NameField
}

func (t *tracingImpl) stopper(err error) {
	t.EndTS = time.Now().UnixMilli()

	// TODO flytta in detta i error handlern (ifen nedan)?
	arr := make([]*types.Stack, 0)

	it := stack.Trace().TrimRuntime()[1:]

	for _, trace := range it {
		frame := trace.Frame()
		arr = append(arr, &types.Stack{
			File: fmt.Sprintf("%+s", trace),
			Line: frame.Line,
			Func: frame.Function,
		})
	}

	if err != nil {
		t.Err = &TracingError{
			Error: err.Error(),
			Stack: arr,
		}
	}

	if t.handler != nil {
		t.handler(t)
	}
}
