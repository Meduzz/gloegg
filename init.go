package gloegg

import (
	"github.com/Meduzz/gloegg/logging"
	"github.com/Meduzz/gloegg/sinks/console"
	"github.com/Meduzz/gloegg/toggles"
)

func init() {
	// enable cosnole output by default
	toggles.SetBoolToggle(console.ConsoleLogEnabled, true)
	// enable log output by default
	toggles.SetBoolToggle(console.ConsolePrintLogEnabled, true)
	// enable trace output by default
	toggles.SetBoolToggle(console.ConsolePrintTraceEnabled, true)

	// add console sink by default
	logging.AddSink(console.NewConsoleWriter())

	// start "sink"-loop
	go logging.StartSink(ingestionChannel, doneChannel)
}
