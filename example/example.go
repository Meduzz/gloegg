package main

import (
	"fmt"
	"strings"

	"github.com/Meduzz/gloegg"
	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/logging"
	"github.com/Meduzz/gloegg/toggles"
)

type (
	GreetingService struct {
		logger common.Logger
	}
)

func NewService() *GreetingService {
	// get or create a new logger called 'GreetingLogic'
	logger := gloegg.CreateLogger("GreetingLogic")

	return &GreetingService{logger}
}

// Greet will greet the provided name as long as it's under 20 characters long
func (g *GreetingService) Greet(name string) (string, error) {
	// create a trace, if ctx contains a traceId it will become the parent of this one
	trace := g.logger.Trace("greeting")

	// lets do a pointless log, that we can also use as metric for length of greeted names
	trace.Info("executing greeting", common.Pair("name", name), common.Pair("length", len(name)))

	// load a feature toggle for the max length of the name parameter
	toggle := toggles.GetIntToggle("name:max.size")

	// traces does logging too, called checkpoints
	trace.Info("about to execute name logic", common.Pair("length", len(name)))
	if len(name) < 10 {
		// close the trace, the nil means no error
		defer trace.Done(nil)

		// good for when need to debug this complicated logic
		trace.Debug("length was under 10")

		return fmt.Sprintf("Hello %s!", strings.ToLower(name)), nil
		// compare len of name to feature toggle for max name length, that defaults to 20
	} else if len(name) > toggle.DefaultValue(20) {
		// good for when need to debug this complicated logic
		trace.Debug("length was more than max allowed")

		err := fmt.Errorf("sorry, your name is too long")
		// turns out there was an error here, we better log that
		trace.Error("length of name was too long", err, common.Pair("actual", len(name)), common.Pair("max", toggle.DefaultValue(20)))

		// close the trace, this time with an error
		trace.Done(err)

		return "", err
	} else {
		// close the trace, without any errors
		defer trace.Done(nil)

		// good for when need to debug this complicated logic
		trace.Debug("length was between 10 & 20")

		return fmt.Sprintf("Hello %s!", strings.ToUpper(name)), nil
	}
}

func main() {
	// set global metadata point service
	gloegg.AddMeta("service", "GreetingService")

	// Setup logger toggle
	toggles.SetStringToggle(logging.FlagForLogger("GreetingLogic"), common.LevelDebug)

	// shut down gloegg in a safe way by draining logs and traces
	defer gloegg.Drain()

	fmt.Println("Enter your name:")
	name := ""
	i, err := fmt.Scanln(&name)

	if err != nil {
		panic(err)
	}

	if i > 0 {
		svc := NewService()
		greeting, err := svc.Greet(name)

		if err != nil {
			panic(err)
		}

		svc.logger.Info(greeting)
	}
}
