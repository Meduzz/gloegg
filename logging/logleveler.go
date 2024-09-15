package logging

import (
	"log/slog"
	"strings"

	"github.com/Meduzz/gloegg/toggles"
)

type (
	logLeveler struct {
		toggle toggles.StringToggle
	}
)

func LevelFromToggle(toggle toggles.StringToggle) slog.Leveler {
	return &logLeveler{toggle}
}

func (l *logLeveler) Level() slog.Level {
	switch strings.ToLower(l.toggle.Value()) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
