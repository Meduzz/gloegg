package toggles

import "github.com/Meduzz/gloegg/common"

type (
	objectToggle struct {
		name     string
		value    map[string]any
		metadata []*common.Tag
	}
)

func newObjectToggle(name string, value map[string]any, metadata []*common.Tag) ObjectToggle {
	return &objectToggle{name, value, metadata}
}

func (o *objectToggle) Matches(other ...*common.Tag) bool {
	return matches(o.metadata, other)
}

func (o *objectToggle) Name() string {
	return o.name
}

func (o *objectToggle) Type() string {
	return KindObject
}

func (o *objectToggle) Value() map[string]any {
	return o.value
}

func (o *objectToggle) Set(v map[string]any) {
	o.value = v
}

func (o *objectToggle) SetField(key string, value any) {
	o.value[key] = value
}

func (o *objectToggle) DefaultInt(key string, defaultValue int) int {
	it, ok := o.value[key]

	if !ok {
		return defaultValue
	}

	value, ok := it.(int)

	if !ok {
		return defaultValue
	}

	return value
}

func (o *objectToggle) DefaultInt64(key string, defaultValue int64) int64 {
	it, ok := o.value[key]

	if !ok {
		return defaultValue
	}

	value, ok := it.(int64)

	if !ok {
		return defaultValue
	}

	return value
}

func (o *objectToggle) DefaultString(key string, defaultValue string) string {
	it, ok := o.value[key]

	if !ok {
		return defaultValue
	}

	value, ok := it.(string)

	if !ok {
		return defaultValue
	}

	return value
}

func (o *objectToggle) DefaultFloat(key string, defaultValue float32) float32 {
	it, ok := o.value[key]

	if !ok {
		return defaultValue
	}

	value, ok := it.(float32)

	if !ok {
		return defaultValue
	}

	return value
}

func (o *objectToggle) DefaultFloat64(key string, defaultValue float64) float64 {
	it, ok := o.value[key]

	if !ok {
		return defaultValue
	}

	value, ok := it.(float64)

	if !ok {
		return defaultValue
	}

	return value
}

func (o *objectToggle) DefaultBool(key string, defaultValue bool) bool {
	it, ok := o.value[key]

	if !ok {
		return defaultValue
	}

	value, ok := it.(bool)

	if !ok {
		return defaultValue
	}

	return value
}
