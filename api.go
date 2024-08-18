package gloegg

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/logging"
	"github.com/Meduzz/gloegg/toggles"
	"github.com/Meduzz/helper/fp/slice"
)

var (
	systemMetadata = common.Tags{}
	handlerFactory = logging.NewTextHandler(os.Stdout)
)

// CreateLogger will create a logger with the provided name
func CreateLogger(name string) *slog.Logger {
	if name == "" {
		return nil
	}

	toggle := toggles.GetStringToggle(fmt.Sprintf("logger.%s", name))
	level := logging.LevelFromToggle(toggle)
	handler := handlerFactory.Spawn(level, systemMetadata)

	return slog.New(handler)
}

// AddMeta add a piece of metadata, removing previous instances of this key
func AddMeta(key string, value any) {
	systemMetadata = slice.Filter(systemMetadata, func(tag *common.Tag) bool {
		return tag.Key != key
	})

	systemMetadata = append(systemMetadata, common.Pair(key, value))
}

// Pair adds another shortcut to create tags
func Pair(key string, value any) *common.Tag {
	return common.Pair(key, value)
}

// Pairs adds another shortcut to creating tags
func Pairs(attrs ...any) []*common.Tag {
	return common.Pairs(attrs...)
}
