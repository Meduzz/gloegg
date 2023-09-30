package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Meduzz/gloegg"
	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/logging"
	"github.com/Meduzz/gloegg/sinks/console"
	"github.com/Meduzz/gloegg/toggles"
	"github.com/Meduzz/gloegg/tracing"
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
func (g *GreetingService) Greet(ctx context.Context, name string) (string, error) {
	// create a trace, if ctx contains a traceId it will become the parent of this one
	trace, err := g.logger.TraceContext("greeting", ctx)

	// ignore unreadable parent trace errors
	if err != nil && !errors.Is(tracing.ErrUnreadableTraceID, err) {
		return "", err
	}

	// lets do a pointless log, that we can also use as metric for length of greeted names
	g.logger.Info("executing greeting", common.Pair("name", name), common.Pair("length", len(name)))

	// load a feature toggle for the max length of the name parameter
	toggle := toggles.GetIntToggle("name:max.size")

	if len(name) < 10 {
		// close the trace, the nil means no error
		defer trace.Done(nil)

		// good for when need to debug this complicated logic
		g.logger.Debug("length was under 10")

		return fmt.Sprintf("Hello %s!", strings.ToLower(name)), nil
		// compare len of name to feature toggle for max name length, that defaults to 20
	} else if len(name) > toggle.DefaultValue(20) {
		toggles.SetBoolToggle(console.ConsolePrintTraceEnabled, true) // print traces

		// good for when need to debug this complicated logic
		g.logger.Debug("length was more than max allowed")

		err := fmt.Errorf("sorry, your name is too long")
		// turns out there was an error here, we better log that
		g.logger.Error("length of name was too long", err, common.Pair("actual", len(name)), common.Pair("max", toggle.DefaultValue(20)))

		// close the trace, this time with an error
		trace.Done(err)

		return "", err
	} else {
		// close the trace, without any errors
		defer trace.Done(nil)

		// good for when need to debug this complicated logic
		g.logger.Debug("length was between 10 & 20")

		return fmt.Sprintf("Hello %s!", strings.ToUpper(name)), nil
	}
}

func main() {
	// set global metadata point service
	gloegg.AddMeta("service", "GreetingService")

	// Setup logger toggle
	settings := toggles.SetObjectToggle(logging.Name("GreetingLogic"), make(map[string]any))
	// enable debug for this logger
	settings.SetField("level", "debug") // default info
	// track all traces
	settings.SetField("tracing", true) // default false

	// toggles.SetBoolToggle(console.ConsoleLogEnabled, false) // disable default console logger
	// toggles.SetBoolToggle(console.ConsoleLogJson, true) // set console logger output to json
	toggles.SetBoolToggle(console.ConsolePrintTraceEnabled, false) // dont print traces

	fmt.Println("Enter your name:")
	name := ""
	i, err := fmt.Scanln(&name)

	if err != nil {
		panic(err)
	}

	if i > 0 {
		svc := NewService()
		greeting, err := svc.Greet(context.Background(), name)

		if err != nil {
			panic(err)
		}

		svc.logger.Info(greeting)
	}

	// shut down gloegg in a safe way by draining logs and traces
	gloegg.Drain()
}
