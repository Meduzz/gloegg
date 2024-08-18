package toggles

type (
	intToggle struct {
		*base
	}
)

func newIntToggle(name string, value int, callback chan *UpdatedToggle) IntToggle {
	return &intToggle{&base{name, KindInt, value, callback}}
}

func (i *intToggle) Value() int {
	result, ok := i.value.(int)

	if !ok {
		return 0
	}

	return result
}

func (i *intToggle) Set(value int) {
	i.value = value
	callbacks <- i.Updated()
}

func (i *intToggle) DefaultValue(value int) int {
	result := i.Value()

	if result == 0 {
		return value
	}

	return result
}

func (i *intToggle) Equals(value int) bool {
	return i.value == value
}

func (i *intToggle) MoreThan(value int) bool {
	return i.Value() > value
}

func (i *intToggle) LessThan(value int) bool {
	return i.Value() < value
}
