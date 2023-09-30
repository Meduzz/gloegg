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
		settings := toggles.GetObjectToggle(Name(event.Logger))

		switch event.Kind {
		case "LOG":
			handleLog(event, settings)
		case "TRACE":
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

func Name(name string) string {
	return fmt.Sprintf("logger.%s", name)
}

func handleLog(event *common.Event, settings toggles.ObjectToggle) {
	logLevel := settings.DefaultString("level", "info")

	if shouldLog(logLevel, event.Log.Level) {
		slice.ForEach(sinks, func(sink common.Sink) {
			sink.Handle(event)
		})
	}
}

func handleTrace(event *common.Event, settings toggles.ObjectToggle) {
	enabled := settings.DefaultBool("tracing", false)

	if enabled {
		slice.ForEach(sinks, func(sink common.Sink) {
			sink.Handle(event)
		})
	}
}

func shouldLog(system, event string) bool {
	allowed := make([]string, 0)

	switch system {
	case "error":
		allowed = append(allowed, "error")
	case "warn":
		allowed = append(allowed, "warn", "error")
	case "info":
		allowed = append(allowed, "info", "warn", "error")
	case "debug":
		allowed = append(allowed, "debug", "info", "error", "warn")
	default:
	}

	return slice.Contains(allowed, event)
}
