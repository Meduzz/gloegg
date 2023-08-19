package toggles

import "github.com/Meduzz/gloegg/common"

type (
	float64Toggle struct {
		name     string
		typ      string
		value    float64
		metadata []*common.Tag
	}
)

func newFloat64Toggle(name string, value float64, metadata []*common.Tag) Float64Toggle {
	return &float64Toggle{name, "float64", value, metadata}
}

func (i *float64Toggle) Matches(needle ...*common.Tag) bool {
	return matches(needle, i.metadata)
}

func (i *float64Toggle) Name() string {
	return i.name
}

func (i *float64Toggle) Type() string {
	return i.typ
}

func (i *float64Toggle) Value() float64 {
	return i.value
}

func (i *float64Toggle) Set(value float64) {
	i.value = value
}

func (i *float64Toggle) DefaultValue(value float64) float64 {
	if i.value == 0 {
		return value
	}

	return i.value
}

func (i *float64Toggle) Equals(value float64) bool {
	return i.value == value
}

func (i *float64Toggle) MoreThan(value float64) bool {
	return i.value > value
}

func (i *float64Toggle) LessThan(value float64) bool {
	return i.value < value
}
