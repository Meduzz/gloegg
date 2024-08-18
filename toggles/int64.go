package toggles

type (
	int64Toggle struct {
		*base
	}
)

func newInt64Toggle(name string, value int64, callback chan *UpdatedToggle) Int64Toggle {
	return &int64Toggle{&base{name, KindInt64, value, callback}}
}

func (i *int64Toggle) Value() int64 {
	result, ok := i.value.(int64)

	if !ok {
		return 0
	}

	return result
}

func (i *int64Toggle) Set(value int64) {
	i.value = value
	i.callback <- i.Updated()
}

func (i *int64Toggle) DefaultValue(value int64) int64 {
	result := i.Value()

	if result == 0 {
		return value
	}

	return result
}

func (i *int64Toggle) Equals(value int64) bool {
	return i.value == value
}

func (i *int64Toggle) MoreThan(value int64) bool {
	return i.Value() > value
}

func (i *int64Toggle) LessThan(value int64) bool {
	return i.Value() < value
}
