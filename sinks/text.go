package sinks

import (
	"fmt"
	"io"

	"github.com/Meduzz/gloegg/common"
)

type (
	textSink struct {
		out    io.Writer
		format common.LogFormat
	}
)

func NewTextSink(out io.Writer, formatter common.LogFormat) common.Sink {
	return &textSink{out, formatter}
}

func (c *textSink) Handle(event *common.Event) {
	_, err := fmt.Fprint(c.out, c.format(event))

	if err != nil {
		println("writing log threw error", err.Error())
	}
}

func (c *textSink) Name() string {
	return "text-sink"
}
