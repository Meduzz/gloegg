package console

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/toggles"
	"github.com/Meduzz/helper/fp/slice"
)

type (
	consoleSink struct{}
)

const (
	ConsoleLogJson           = "gloegg.log.json"
	ConsoleLogEnabled        = "gloegg.log.enabled"
	ConsolePrintLogEnabled   = "gloegg.print.log.enabled"
	ConsolePrintTraceEnabled = "gloegg.print.trace.enabled"
)

func NewConsoleWriter() common.Sink {
	return &consoleSink{}
}

func (c *consoleSink) Handle(event *common.Event) {
	enabledFlag := toggles.GetBoolToggle(ConsoleLogEnabled)
	jsonFlag := toggles.GetBoolToggle(ConsoleLogJson)

	if enabledFlag.Value() {
		if jsonFlag.Value() {
			bs, _ := json.Marshal(event)

			fmt.Println(string(bs))
		} else {
			showLogs := toggles.GetBoolToggle(ConsolePrintLogEnabled)
			showTraces := toggles.GetBoolToggle(ConsolePrintTraceEnabled)

			created := event.Created.Format(time.DateTime)
			logger := event.Logger

			if event.Kind == "LOG" && showLogs.Value() {
				// ts [logger] level - message (metadata)
				level := event.Log.Level

				buf := bytes.NewBufferString("")
				fmt.Fprintf(buf, "%s [%s] %s - %s [%s]\n", created, logger, level, event.Log.Message, dumpMetadata(event.Metadata))

				if event.Log.Error != "" {
					// <error message>
					fmt.Fprintf(buf, "Error:\n\t%s\n", event.Log.Error)
					fmt.Fprintln(buf, "Stacktrace:")

					// func - file:line
					slice.ForEach(event.Log.StackTrace, func(item *common.Stack) {
						// TODO filter external toggle, that prefix check the item.Func to get rid of noise?
						fmt.Fprintf(buf, "\t%s - %s:%d\n", item.Func, item.File, item.Line)
					})
				}

				fmt.Print(buf.String())
			} else if event.Kind == "TRACE" && showTraces.Value() {
				// ts [logger] trace: message (duration?) (metadata)
				start := event.Trace.Start
				end := event.Trace.End

				fmt.Printf("%s [%s] trace: %s (%s) [%s]\n", created, logger, event.Trace.Name, end.Sub(start).String(), dumpMetadata(event.Metadata))

				if showLogs.Value() {
					slice.ForEach(event.Trace.Checkpoints, func(it *common.CheckpointDTO) {
						// ts [logger] checkpoint - message (duration) metadata

						buf := bytes.NewBufferString("")
						fmt.Fprintf(buf, "%s [%s] checkpoint (%s) - %s (%s) [%s]\n", it.Created.Format(time.DateTime), logger, it.Level, it.Message, it.Created.Sub(event.Trace.Start), dumpMetadata(it.Metadata))

						if it.Error != "" {
							// <error message>
							fmt.Fprintf(buf, "Error:\n\t%s\n", it.Error)
							fmt.Fprintln(buf, "Stacktrace:")

							// func - file:line
							slice.ForEach(it.StackTrace, func(item *common.Stack) {
								// TODO filter external toggle, that prefix check the item.Func to get rid of noise?
								fmt.Fprintf(buf, "\t%s - %s:%d\n", item.Func, item.File, item.Line)
							})
						}

						fmt.Print(buf.String())
					})
				}
			}
		}
	}
}

func dumpMetadata(tags []*common.Tag) string {
	converted := slice.Map(tags, func(tag *common.Tag) string {
		return fmt.Sprintf("%s=%v", tag.Key, tag.Value)
	})

	return strings.Join(converted, ", ")
}
