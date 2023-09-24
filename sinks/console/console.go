package console

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/toggles"
	"github.com/Meduzz/helper/fp/slice"
)

type (
	consoleSink struct{}
)

const (
	ConsoleLogJson    = "gloegg.log.json"
	ConsoleLogEnabled = "gloegg.log.enabled"
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
			created := event.Created.Format(time.DateTime)
			logger := event.Logger

			if event.Kind == "LOG" {
				// ts [logger] level - message (metadata)
				level := event.Log.Level

				buf := bytes.NewBufferString("")
				fmt.Fprintf(buf, "%s [%s] %s - %s metadata{%v}\n", created, logger, level, event.Log.Message, common.AsMap(event.Metadata...))

				if event.Log.Error != nil {
					// <error message>
					fmt.Fprintf(buf, "Error message: %s\n", event.Log.Error.Error())

					// func - file:line
					slice.ForEach(event.Log.StackTrace, func(item *common.Stack) {
						// TODO filter external toggle, that prefix check the item.Func to get rid of noise?
						fmt.Fprintf(buf, "\t%s - %s:%d\n", item.Func, item.File, item.Line)
					})
				}

				fmt.Print(buf.String())
			} else if event.Kind == "TRACE" {
				// ts [logger] - message (duration?) (metadata)
				start := event.Trace.Start
				end := event.Trace.End

				fmt.Printf("%s [%s] - %s (%s) metadata{%v}\n", created, logger, event.Trace.Name, end.Sub(start).String(), common.AsMap(event.Metadata...))
			}
		}
	}
}
