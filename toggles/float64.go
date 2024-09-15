package toggles

type (
	float64Toggle struct {
		*base
	}
)

func newFloat64Toggle(name string, value float64, callback chan *UpdatedToggle) Float64Toggle {
	return &float64Toggle{&base{name, KindFloat64, value, callback}}
}

func (i *float64Toggle) Value() float64 {
	result, ok := i.value.(float64)

	if !ok {
		return 0
	}

	return result
}

func (i *float64Toggle) Set(value float64) {
	i.value = value
	i.callback <- i.Updated()
}

func (i *float64Toggle) DefaultValue(value float64) float64 {
	result := i.Value()

	if result == 0 {
		return value
	}

	return result
}

func (i *float64Toggle) Equals(value float64) bool {
	return i.value == value
}

func (i *float64Toggle) MoreThan(value float64) bool {
	return i.Value() > value
}

func (i *float64Toggle) LessThan(value float64) bool {
	return i.Value() < value
}
