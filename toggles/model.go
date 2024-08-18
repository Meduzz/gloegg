package toggles

type (
	Toggle interface {
		Name() string
		Kind() string
	}

	StringToggle interface {
		Toggle
		Value() string
		Set(string)
		DefaultValue(string) string
		Equals(string) bool
		Contains(string) bool
	}

	IntToggle interface {
		Toggle
		Value() int
		Set(int)
		DefaultValue(int) int
		Equals(int) bool
		MoreThan(int) bool
		LessThan(int) bool
	}

	Int64Toggle interface {
		Toggle
		Value() int64
		Set(int64)
		DefaultValue(int64) int64
		Equals(int64) bool
		MoreThan(int64) bool
		LessThan(int64) bool
	}

	Float64Toggle interface {
		Toggle
		Value() float64
		Set(float64)
		DefaultValue(float64) float64
		Equals(float64) bool
		MoreThan(float64) bool
		LessThan(float64) bool
	}

	Float32Toggle interface {
		Toggle
		Value() float32
		Set(float32)
		DefaultValue(float32) float32
		Equals(float32) bool
		MoreThan(float32) bool
		LessThan(float32) bool
	}

	BoolToggle interface {
		Toggle
		Value() bool
		Set(bool)
	}

	ObjectToggle interface {
		Toggle
		Value() map[string]any
		Set(map[string]any)
		SetField(string, any)
		DefaultInt(string, int) int
		DefaultInt64(string, int64) int64
		DefaultString(string, string) string
		DefaultFloat(string, float32) float32
		DefaultFloat64(string, float64) float64
		DefaultBool(string, bool) bool
	}

	base struct {
		name     string
		kind     string
		value    any
		callback chan *UpdatedToggle
	}

	UpdatedToggle struct {
		Name  string
		Kind  string
		Value any
	}
)

func (b *base) Name() string {
	return b.name
}

func (b *base) Kind() string {
	return b.kind
}

func (b *base) Updated() *UpdatedToggle {
	return &UpdatedToggle{
		Name:  b.name,
		Kind:  b.kind,
		Value: b.value,
	}
}
