# gloegg (glögg)

Gloegg has been on a diet. Left are the tags and the feature toggles. Logging has been replaced by `log/slog`. Tracing are nowhere to be found. Maybe it will return in some form or shape one day.

## Todo

* Allow to set handlerFactory in root (gloegg)
    * On that note, create another handlerFactory for json
* Prolly hide most of the internal stuff in internal folders
* Formalize how to subscribe to toggle updates
    * On that note, that channel might cause problems for people unless addressed