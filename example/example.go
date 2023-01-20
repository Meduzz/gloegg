package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Meduzz/gloegg"
	"github.com/Meduzz/gloegg/toggles"
	"github.com/Meduzz/gloegg/types"
	"github.com/rs/zerolog"
)

type (
	GreetingService struct {
		logger types.Logging
	}
)

func NewService() *GreetingService {
	// get or create a new logger called 'GreetingService'
	logger := gloegg.Logger("greetingService.GreetingService")
	return &GreetingService{logger}
}

// Greet will greet the provided name as long as it's under 20 characters long
func (g *GreetingService) Greet(ctx context.Context, name string) (string, error) {
	// create a trace, if ctx contains a traceId it will become the parent of this one
	_, done := g.logger.TraceContext("greeting", ctx)

	// lets do a pointless log, that we can also use as metric for length of greeted names
	g.logger.Info("executing greeting", types.Pair("name", name), types.Pair("length", len(name)))

	// load a feature toggle for the max length of the name parameter
	toggle := g.logger.FeatureToggle("name:max.size")

	if len(name) < 10 {
		// good for when need to debug this complicated logic
		g.logger.Debug("length was under 10")

		// close the trace, the nil means no error
		defer done(nil)
		return fmt.Sprintf("Hello %s!", strings.ToLower(name)), nil
		// compare len of name to feature toggle for max name length, that defaults to 20
	} else if len(name) > toggle.GetInt(20) {
		// good for when need to debug this complicated logic
		g.logger.Debug("length was more than max allowed")

		err := fmt.Errorf("sorry, your name is too long")
		// turns out there was an error here, we better log that
		g.logger.Error("length of name", err, types.Pair("actual", len(name)), types.Pair("max", toggle.GetInt(20)))

		// close the trace, this time with an error
		defer done(err)
		return "", err
	} else {
		// good for when need to debug this complicated logic
		g.logger.Debug("length was between 10 & 20")

		// close the trace, without any errors
		defer done(nil)
		return fmt.Sprintf("Hello %s!", strings.ToUpper(name)), nil
	}
}

func main() {
	// Remove me for "vanilla" feel
	setupLogging()

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

		fmt.Println(greeting)
	}

	fmt.Println("Done!")
}

func setupLogging() {
	// call setup logging to setup root logger
	gloegg.SetupLogging(func() zerolog.Logger {
		// use glögg defaults (time=unixms & generate stacktraces)
		gloegg.LoggingDefaults()

		// FYI setting global logging level will override any logging level you set with toggles :(

		// set debug as default level for greetingService.GreetingService logger.
		toggles.SetToggle("greetingService.GreetingService", "debug")

		// create and setup the actual logger
		return zerolog.New(os.Stdout).With().Stack().Timestamp().Str("service", "GreetingService").Logger()
	})

	logger := gloegg.Logger("traces")

	// set trace handler, that logs traces
	gloegg.SetTraceHandler(traceHandler(logger))
}

func traceHandler(l types.Logging) func(types.Trace) {
	return func(t types.Trace) {
		bs, err := json.Marshal(t)

		if err == nil {
			l.Info("trace complete", types.Pair("trace", json.RawMessage(bs)))
		}
	}
}
