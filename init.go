package gloegg

import (
	"os"

	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/logging"
	"github.com/Meduzz/gloegg/sinks"
)

func init() {
	// add console sink by default
	logging.AddSink(sinks.NewTextSink(os.Stderr, common.DefaultFormat))

	// start "sink"-loop
	go logging.StartSink(ingestionChannel, doneChannel)
}
