package toggles

import "github.com/Meduzz/gloegg/common"

type (
	intToggle struct {
		name     string
		typ      string
		value    int
		metadata []*common.Tag
	}
)

func newIntToggle(name string, value int, metadata []*common.Tag) IntToggle {
	return &intToggle{name, "int", value, metadata}
}

func (i *intToggle) Matches(needle ...*common.Tag) bool {
	return matches(needle, i.metadata)
}

func (i *intToggle) Name() string {
	return i.name
}

func (i *intToggle) Type() string {
	return i.typ
}

func (i *intToggle) Value() int {
	return i.value
}

func (i *intToggle) Set(value int) {
	i.value = value
}

func (i *intToggle) DefaultValue(value int) int {
	if i.value == 0 {
		return value
	}

	return i.value
}

func (i *intToggle) Equals(value int) bool {
	return i.value == value
}

func (i *intToggle) MoreThan(value int) bool {
	return i.value > value
}

func (i *intToggle) LessThan(value int) bool {
	return i.value < value
}
