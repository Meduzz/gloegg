package toggles

type (
	float32Toggle struct {
		*base
	}
)

func newFloat32Toggle(name string, value float32, callbacks chan *UpdatedToggle) Float32Toggle {
	return &float32Toggle{&base{name, KindFloat32, value, callbacks}}
}

func (i *float32Toggle) Value() float32 {
	result, ok := i.value.(float32)

	if !ok {
		return 0
	}

	return result
}

func (i *float32Toggle) Set(value float32) {
	i.value = value
	i.callback <- i.Updated()
}

func (i *float32Toggle) DefaultValue(value float32) float32 {
	result := i.Value()

	if result == 0 {
		return value
	}

	return result
}

func (i *float32Toggle) Equals(value float32) bool {
	return i.value == value
}

func (i *float32Toggle) MoreThan(value float32) bool {
	return i.Value() > value
}

func (i *float32Toggle) LessThan(value float32) bool {
	return i.Value() < value
}
