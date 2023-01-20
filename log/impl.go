package log

import (
	"context"

	"github.com/Meduzz/gloegg/toggles"
	"github.com/Meduzz/gloegg/tracing"
	"github.com/Meduzz/gloegg/types"
	"github.com/rs/zerolog"
)

type (
	logger struct {
		ctx          context.Context
		log          zerolog.Logger
		traceHandler func(types.Trace)
	}
)

func New(log zerolog.Logger, traceHandler func(types.Trace)) types.Logging {
	return &logger{
		ctx:          context.Background(),
		log:          log,
		traceHandler: traceHandler,
	}
}

func From(ctx context.Context, log zerolog.Logger, traceHandler func(types.Trace)) types.Logging {
	return &logger{
		ctx:          ctx,
		log:          log,
		traceHandler: traceHandler,
	}
}

func (l *logger) Info(msg string, tags ...*types.Tag) {
	e := l.log.Info()

	for _, tag := range tags {
		e.Interface(tag.Key, tag.Value)
	}

	e.Msg(msg)
}

func (l *logger) Debug(msg string, tags ...*types.Tag) {
	e := l.log.Debug()

	for _, tag := range tags {
		e.Interface(tag.Key, tag.Value)
	}

	e.Msg(msg)
}

func (l *logger) Warn(msg string, tags ...*types.Tag) {
	e := l.log.Warn()

	for _, tag := range tags {
		e.Interface(tag.Key, tag.Value)
	}

	e.Msg(msg)
}

func (l *logger) Error(msg string, err error, tags ...*types.Tag) {
	e := l.log.Error()

	for _, tag := range tags {
		e.Interface(tag.Key, tag.Value)
	}

	e.Err(err).Msg(msg)
}

func (l *logger) FeatureToggle(name string, tags ...*types.Tag) types.Toggle {
	t := toggles.GetToggle(name, tags...)

	if t == nil {
		t = toggles.SetToggle(name, nil, tags...)
	}

	return t
}

func (l *logger) Trace(name string, tags ...*types.Tag) (context.Context, func(error)) {
	trace := tracing.New(name, l.ctx, l.traceHandler, tags...)

	return trace.Start()
}

func (l *logger) TraceContext(name string, parent context.Context, tags ...*types.Tag) (context.Context, func(error)) {
	trace := tracing.New(name, parent, l.traceHandler, tags...)

	return trace.Start()
}
