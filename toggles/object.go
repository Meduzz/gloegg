package toggles

type (
	objectToggle struct {
		*base
	}
)

func newObjectToggle(name string, value map[string]any, callback chan *UpdatedToggle) ObjectToggle {
	return &objectToggle{&base{name, KindObject, value, callback}}
}

func (o *objectToggle) Value() map[string]any {
	result, ok := o.value.(map[string]any)

	if !ok {
		return make(map[string]any)
	}

	return result
}

func (o *objectToggle) Set(v map[string]any) {
	o.value = v
	o.callback <- o.Updated()
}

func (o *objectToggle) SetField(key string, value any) {
	result, ok := o.value.(map[string]any)

	if !ok {
		return
	}

	result[key] = value
}

func (o *objectToggle) DefaultInt(key string, defaultValue int) int {
	result, ok := o.value.(map[string]any)

	if !ok {
		return defaultValue
	}

	it, ok := result[key]

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
	result, ok := o.value.(map[string]any)

	if !ok {
		return defaultValue
	}

	it, ok := result[key]

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
	result, ok := o.value.(map[string]any)

	if !ok {
		return defaultValue
	}

	it, ok := result[key]

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
	result, ok := o.value.(map[string]any)

	if !ok {
		return defaultValue
	}

	it, ok := result[key]

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
	result, ok := o.value.(map[string]any)

	if !ok {
		return defaultValue
	}

	it, ok := result[key]

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
	result, ok := o.value.(map[string]any)

	if !ok {
		return defaultValue
	}

	it, ok := result[key]

	if !ok {
		return defaultValue
	}

	value, ok := it.(bool)

	if !ok {
		return defaultValue
	}

	return value
}
