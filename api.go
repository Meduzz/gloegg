package gloegg

import (
	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/gloegg/logging"
	"github.com/Meduzz/helper/fp/slice"
)

var (
	systemMetadata   = make([]*common.Tag, 0)
	ingestionChannel = make(chan *common.Event)
	doneChannel      = make(chan int, 1)
)

// CreateLogger will create a logger with the provided name
func CreateLogger(name string) common.Logger {
	if name == "" {
		return nil
	}

	return logging.NewLogger(name, ingestionChannel, systemMetadata)
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

// Drain logs and traces then close down gloegg
func Drain() {
	close(ingestionChannel)
	<-doneChannel
}
