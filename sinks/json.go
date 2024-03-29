package sinks

import (
	"encoding/json"
	"io"

	"github.com/Meduzz/gloegg/common"
)

type (
	jsonSink struct {
		out io.Writer
	}
)

func NewJsonSink(out io.Writer) common.Sink {
	return &jsonSink{out}
}

func (j *jsonSink) Handle(event *common.Event) {
	bs, err := json.Marshal(event)

	if err != nil {
		println("marshal even to json threw error", err.Error())
		return
	}

	_, err = j.out.Write(bs)

	if err != nil {
		println("writing json log threw error", err.Error())
	}
}

func (j *jsonSink) Name() string {
	return "json-sink"
}
