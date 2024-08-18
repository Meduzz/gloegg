package toggles

type (
	boolToggle struct {
		*base
	}
)

func newBoolToggle(name string, value bool, callbacks chan *UpdatedToggle) BoolToggle {
	return &boolToggle{&base{name, KindBool, value, callbacks}}
}

func (b *boolToggle) Value() bool {
	result, ok := b.value.(bool)

	if !ok {
		return false
	}

	return result
}

func (b *boolToggle) Set(value bool) {
	b.value = value
	b.callback <- b.Updated()
}
