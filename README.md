# gloegg (glögg)
tacos made up of logging, metrics, traces and feature toggles. In general these "tools" are not in the same package, but I think they relate... one way or the other.

Throughout the lib, there's places where tags can be used to drill down into the data once it's been collected.

## Loggers
The loggers is where you "collect" all your data. This is where you enter log-lines. As well as starting traces. A logger is sort of a scope to the data it generates. Loggers can have individual settings, like level and tracing on/off.

The level of the logger defines which kinds of log events that gets forwarded to the sinks. Currently these levels are defined:

* debug
* info
* warn
* error

Setting the level of a logger to warn, would only forward events logged in level `warn` or `error`.

Currently traces for a logger are either on or off. With tracing on, all traces from the logger are forwarded to the sinks.

## Sinks
Eventually, if all conditions are met, your log or trace event ends up in a sink. Currently there's only one default sink, it spews out the data into stdout (in a semi crappy format). It has a few bool-toggles that can be tweaked, see consts in `sinks/console` package or `example/example.go` for usage.

You can create and add your own sinks.

### ConsoleSink
"Bundled" with the lib comes a sink that will write your logs and traces to stdout. It relies heavily on feature toggles though. It can be fully disabled by setting boolean toogle `gloegg.log.enabled`. It can be instructed to generate json logging via boolean toggle `gloegg.log.json`. Printing of logs and traces individually can be controlled with boolean toggles `gloegg.print.log.enabled` (logs (& checkpoints)) & `gloegg.print.trace.enabled` (traces).

While default values of boolean toggles are `false`, logs and traces (with checkpoints) will be enabled explicitly by the lib (but can be disabled via the above flags).

## Logging
This part does structured logging, pretty much as expected. The thought is to keep the log line static and add metadata in tags until it makes sense. That metadata could also be metrics, coincidently the next topic.

## Metrics
This part is sort of hidden. Add it to the logs, and use the logs as the name of the metric. Push the complexity with metric type (histogram, gauge, etc) to a server/the cloud. Treat your service telemetry like some stupid IoT sensor data.

## Traces
This part is pretty much a "batteries excluded" solution ;) There's infrastructure to generate simple traces, and a way to collect them. But they will not integrate with any standard systems as of now.

### Checkpoints
Traces can also take logs, but it will as a bonus keep track on when they were generated. So a duration can be calculated against the events created timestamp. Checkpoints looks and feels like logs (they even follow the same rules (mostly)), but are called checkpoints.

## Feature toggles
These got in here while I had the steam up. I wanted a way to control my loggers, which is traditionally done through configs. Which in my mind is narrow minded. Again, it's a bit of a batteries excluded situation. However there should be enough infrastructure to make things spinning, or what ever you'd like them to do.

## Mental picture

When the code runs, it's in silly mode. It does nothing (but waiting for log events). The default console-logger is attached, and it is enabled.

To get things spinning, you need a logger. They must have a name, and it's through them you can do things, like logging and start traces. You get a logger via `gloegg.CreateLogger`-method. Remember, all loggers must have a name!

You can fetch a toggle via any of the `toggles.Get<type>Toggle(name string)`. If the logger has not been created yet, you need to call `toggles.Set<type>Toggle(name string, value <type>)`. With the toggle you can start to read and tweak things.

Loggers use `ObjectLoggers`, these wrap a simple map ex: `settings := toggles.GetObjectToggle("logger.MyLogger")`. Log level is a string, and is changed via `setting.SetField("level", "<new log level [info,error,warn,debug (anything else disalbes the logger)]>")`. While to enable/disable tracing you set `setting.SetField("tracing", <enable/disable tracing [true|false]>)`.
