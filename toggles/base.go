package toggles

import (
	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/helper/fp/slice"
)

var (
	typedToggles []Toggle
)

func SetStringToggle(name, value string, selectors ...*common.Tag) StringToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "string" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(StringToggle)

		if !ok {
			t = newStringToggle(name, value, selectors)
		} else {
			t.Set(value)
		}

		return t
	} else {
		t := newStringToggle(name, value, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func SetIntToggle(name string, value int, selectors ...*common.Tag) IntToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "int" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(IntToggle)

		if !ok {
			t = newIntToggle(name, value, selectors)
		} else {
			t.Set(value)
		}

		return t
	} else {
		t := newIntToggle(name, value, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func SetInt64Toggle(name string, value int64, selectors ...*common.Tag) Int64Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "int64" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Int64Toggle)

		if !ok {
			t = newInt64Toggle(name, value, selectors)
		} else {
			t.Set(value)
		}

		return t
	} else {
		t := newInt64Toggle(name, value, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func SetFloat64Toggle(name string, value float64, selectors ...*common.Tag) Float64Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "float64" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Float64Toggle)

		if !ok {
			t = newFloat64Toggle(name, value, selectors)
		} else {
			t.Set(value)
		}

		return t
	} else {
		t := newFloat64Toggle(name, value, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func SetFloat32Toggle(name string, value float32, selectors ...*common.Tag) Float32Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "float32" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Float32Toggle)

		if !ok {
			t = newFloat32Toggle(name, value, selectors)
		} else {
			t.Set(value)
		}

		return t
	} else {
		t := newFloat32Toggle(name, value, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func SetBoolToggle(name string, value bool, selectors ...*common.Tag) BoolToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "bool" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(BoolToggle)

		if !ok {
			t = newBoolToggle(name, value, selectors)
		} else {
			t.Set(value)
		}

		return t
	} else {
		t := newBoolToggle(name, value, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func SetObjectToggle(name string, value map[string]any, selectors ...*common.Tag) ObjectToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "object" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(ObjectToggle)

		if !ok {
			t = newObjectToggle(name, value, selectors)
		} else {
			t.Set(value)
		}

		return t
	} else {
		t := newObjectToggle(name, value, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetStringToggle(name string, selectors ...*common.Tag) StringToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "string" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(StringToggle)

		if !ok {
			t = newStringToggle(name, "", selectors)
		}

		return t
	}

	return newStringToggle(name, "", selectors)
}

func GetIntToggle(name string, selectors ...*common.Tag) IntToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "int" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(IntToggle)

		if !ok {
			t = newIntToggle(name, 0, selectors)
		}

		return t
	}

	return newIntToggle(name, 0, selectors)
}

func GetInt64Toggle(name string, selectors ...*common.Tag) Int64Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "int64" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Int64Toggle)

		if !ok {
			t = newInt64Toggle(name, 0, selectors)
		}

		return t
	}

	return newInt64Toggle(name, 0, selectors)
}

func GetFloat64Toggle(name string, selectors ...*common.Tag) Float64Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "float64" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Float64Toggle)

		if !ok {
			t = newFloat64Toggle(name, 0, selectors)
		}

		return t
	}

	return newFloat64Toggle(name, 0, selectors)
}

func GetFloat32Toggle(name string, selectors ...*common.Tag) Float32Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "float32" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Float32Toggle)

		if !ok {
			t = newFloat32Toggle(name, 0, selectors)
		}

		return t
	}

	return newFloat32Toggle(name, 0, selectors)
}

func GetBoolToggle(name string, selectors ...*common.Tag) BoolToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "bool" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(BoolToggle)

		if !ok {
			t = newBoolToggle(name, false, selectors)
		}

		return t
	}

	return newBoolToggle(name, false, selectors)
}

func GetObjectToggle(name string, selectors ...*common.Tag) ObjectToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == "object" && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(ObjectToggle)

		if !ok {
			t = newObjectToggle(name, make(map[string]any), selectors)
		}

		return t
	}

	return newObjectToggle(name, make(map[string]any), selectors)
}
