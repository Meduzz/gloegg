package toggles

import "github.com/Meduzz/gloegg/common"

type (
	float32Toggle struct {
		name     string
		typ      string
		value    float32
		metadata []*common.Tag
	}
)

func newFloat32Toggle(name string, value float32, metadata []*common.Tag) Float32Toggle {
	return &float32Toggle{name, "float32", value, metadata}
}

func (i *float32Toggle) Matches(needle ...*common.Tag) bool {
	return matches(needle, i.metadata)
}

func (i *float32Toggle) Name() string {
	return i.name
}

func (i *float32Toggle) Type() string {
	return i.typ
}

func (i *float32Toggle) Value() float32 {
	return i.value
}

func (i *float32Toggle) Set(value float32) {
	i.value = value
}

func (i *float32Toggle) DefaultValue(value float32) float32 {
	if i.value == 0 {
		return value
	}

	return i.value
}

func (i *float32Toggle) Equals(value float32) bool {
	return i.value == value
}

func (i *float32Toggle) MoreThan(value float32) bool {
	return i.value > value
}

func (i *float32Toggle) LessThan(value float32) bool {
	return i.value < value
}
