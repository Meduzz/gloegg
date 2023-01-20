# gloegg (glögg)
tacos made up of logging, metrics, traces and feature toggles. In general these "tools" are in the same package, because they relate... one way or the other.

Though out the lib, there's places where tags can be used to drill down into the data once it's been collected.

## Logging
Does structured logging, pretty much as expected. It's a limited wrapper of the excellent zerolog library. One huge benefit of zerolog is that it can log as JSON, which as you know open doors for automation and analyses etc.

## Metrics
Is hidden, sort of. Add it to the logs, and use the logs as the name of the metric. Push the complexity with metric type (histogram, gauge, etc) to a server/the cloud.

## Traces
This is sort of a batteries excluded solution ;) There's infrastructure to generate simple traces, and a way to collect them. But they will not integrate with any standard systems as of now.

## Feature toggles
These got in here while I had the steam up. I wanted a way to control my loggers, which is traditionally done through configs. Which in my mind is narrow minded. Again, it's a bit of a batteries excluded situation. However there should be enough infrastructure to make things spinning, or what ever you'd like them to do.