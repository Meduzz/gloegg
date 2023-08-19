package toggles

import "github.com/Meduzz/gloegg/common"

type (
	int64Toggle struct {
		name     string
		typ      string
		value    int64
		metadata []*common.Tag
	}
)

func newInt64Toggle(name string, value int64, metadata []*common.Tag) Int64Toggle {
	return &int64Toggle{name, "int64", value, metadata}
}

func (i *int64Toggle) Matches(needle ...*common.Tag) bool {
	return matches(needle, i.metadata)
}

func (i *int64Toggle) Name() string {
	return i.name
}

func (i *int64Toggle) Type() string {
	return i.typ
}

func (i *int64Toggle) Value() int64 {
	return i.value
}

func (i *int64Toggle) Set(value int64) {
	i.value = value
}

func (i *int64Toggle) DefaultValue(value int64) int64 {
	if i.value == 0 {
		return value
	}

	return i.value
}

func (i *int64Toggle) Equals(value int64) bool {
	return i.value == value
}

func (i *int64Toggle) MoreThan(value int64) bool {
	return i.value > value
}

func (i *int64Toggle) LessThan(value int64) bool {
	return i.value < value
}
