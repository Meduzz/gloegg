package main

import (
	"github.com/Meduzz/gloegg"
	"github.com/Meduzz/gloegg/toggles"
)

func main() {
	gloegg.AddMeta("app", "example")
	logger := gloegg.CreateLogger("bengt")

	logger.Info("Bengt happened", "story", true)

	toggle := toggles.GetStringToggle("logger.bengt")
	toggle.Set("warn")

	logger.Info("hidden messages", "in", "plain sight") // not shown...
	logger.Warn("pikaboo")
}
