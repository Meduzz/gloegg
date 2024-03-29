package toggles

import "github.com/Meduzz/gloegg/common"

type (
	boolToggle struct {
		name     string
		value    bool
		metadata []*common.Tag
	}
)

func newBoolToggle(name string, value bool, metadata []*common.Tag) BoolToggle {
	return &boolToggle{name, value, metadata}
}

func (b *boolToggle) Matches(needle ...*common.Tag) bool {
	return matches(needle, b.metadata)
}

func (b *boolToggle) Name() string {
	return b.name
}

func (b *boolToggle) Type() string {
	return KindBool
}

func (b *boolToggle) Value() bool {
	return b.value
}

func (b *boolToggle) Set(value bool) {
	b.value = value
}
