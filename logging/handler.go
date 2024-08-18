package logging

import (
	"io"
	"log/slog"

	"github.com/Meduzz/gloegg/common"
)

type (
	HandlerFactory interface {
		Spawn(slog.Leveler, common.Tags) slog.Handler
	}

	textHandlerFactory struct {
		writer io.Writer
	}
)

func NewTextHandler(writer io.Writer) HandlerFactory {
	return &textHandlerFactory{writer}
}

func (l *textHandlerFactory) Spawn(logLevel slog.Leveler, tags common.Tags) slog.Handler {
	return slog.NewTextHandler(l.writer, &slog.HandlerOptions{Level: logLevel}).WithAttrs(tags.ToSlog())
}
