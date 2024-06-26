package logging

import (
	"fmt"

	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/toggles"
	"github.com/Meduzz/helper/fp/slice"
)

var (
	sinks = make([]common.Sink, 0)
)

func StartSink(channel chan *common.Event, doneChan chan int) {
	for event := range channel {
		settings := toggles.GetStringToggle(FlagForLogger(event.Logger))

		switch event.Kind {
		case common.KindLog:
			handleLog(event, settings)
		case common.KindTrace:
			handleTrace(event, settings)
		default:
			fmt.Printf("unknown kind: %s, dropping\n", event.Kind)
		}
	}

	doneChan <- 1
}

func AddSink(sink common.Sink) {
	sinks = append(sinks, sink)
}

func RemoveSink(sink common.Sink) {
	sinks = slice.Filter(sinks, func(s common.Sink) bool {
		return s.Name() != sink.Name()
	})
}

func FlagForLogger(name string) string {
	return fmt.Sprintf("logger.%s", name)
}

func handleLog(event *common.Event, settings toggles.StringToggle) {
	logLevel := settings.DefaultValue("info")

	if shouldLog(logLevel, event.Log.Level) {
		slice.ForEach(sinks, func(sink common.Sink) {
			sink.Handle(event)
		})
	}
}

func handleTrace(event *common.Event, settings toggles.StringToggle) {
	logLevel := settings.DefaultValue("info")

	event.Trace.Checkpoints = slice.Filter(event.Trace.Checkpoints, func(log *common.CheckpointDTO) bool {
		return shouldLog(logLevel, log.Level)
	})

	slice.ForEach(sinks, func(sink common.Sink) {
		sink.Handle(event)
	})
}

func shouldLog(system, event string) bool {
	allowed := make([]string, 0)

	switch system {
	case common.LevelError:
		allowed = append(allowed, common.LevelError)
	case common.LevelWarn:
		allowed = append(allowed, common.LevelWarn, common.LevelError)
	case common.LevelInfo:
		allowed = append(allowed, common.LevelInfo, common.LevelWarn, common.LevelError)
	case common.LevelDebug:
		allowed = append(allowed, common.LevelDebug, common.LevelInfo, common.LevelWarn, common.LevelError)
	default:
	}

	return slice.Contains(allowed, event)
}
