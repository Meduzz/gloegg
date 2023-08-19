package logging

import (
	"fmt"
	"testing"

	"github.com/Meduzz/gloegg/common"
)

func TestLogger(t *testing.T) {
	channel := make(chan *common.Event, 10)
	metadata := make([]*common.Tag, 0)

	subject := &loggingFacade{channel, metadata, ""}

	trace := subject.Trace("TestLogging", common.Pair("trace", 1))

	meta := make([]*common.Tag, 0)
	meta = append(meta, common.Pair("key1", "value1"))

	subject.Info("info", meta...)
	assertLog(t, channel, "info", "info", meta, nil)

	subject.Debug("debug", meta...)
	assertLog(t, channel, "debug", "debug", meta, nil)

	subject.Warn("warn", meta...)
	assertLog(t, channel, "warn", "warn", meta, nil)

	subject.Error("error", fmt.Errorf("booh"), meta...)
	assertLog(t, channel, "error", "error", meta, fmt.Errorf("booh"))

	trace.Done(nil)
}

func assertLog(t *testing.T, channel chan *common.Event, msg, level string, metadata []*common.Tag, err error) {
	event := <-channel

	if event.Log.Message != msg {
		t.Errorf("message was not the expected, was %s", msg)
	}

	if len(event.Metadata) != len(metadata) {
		t.Errorf("number of metadata was not the expected, was %d", len(event.Metadata))
	}

	if event.Log.Level != level {
		t.Errorf("level was not the expected, was %s", event.Log.Level)
	}

	if err != nil {
		if event.Log.Error == nil {
			t.Error("expected an error")
		}
	}
}
