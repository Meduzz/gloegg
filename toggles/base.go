package toggles

import (
	"fmt"

	"github.com/Meduzz/gloegg/common"
	"github.com/Meduzz/helper/fp/slice"
)

var (
	typedToggles []Toggle
)

const (
	KindString  = "string"
	KindInt     = "int"
	KindInt64   = "int64"
	KindFloat64 = "float64"
	KindFloat32 = "float32"
	KindBool    = "bool"
	KindObject  = "object"
)

func SetStringToggle(name, value string, selectors ...*common.Tag) StringToggle {
	toggle := GetStringToggle(name, selectors...)
	toggle.Set(value)

	return toggle
}

func SetIntToggle(name string, value int, selectors ...*common.Tag) IntToggle {
	toggle := GetIntToggle(name, selectors...)
	toggle.Set(value)

	return toggle
}

func SetInt64Toggle(name string, value int64, selectors ...*common.Tag) Int64Toggle {
	toggle := GetInt64Toggle(name, selectors...)
	toggle.Set(value)

	return toggle
}

func SetFloat64Toggle(name string, value float64, selectors ...*common.Tag) Float64Toggle {
	toggle := GetFloat64Toggle(name, selectors...)
	toggle.Set(value)

	return toggle
}

func SetFloat32Toggle(name string, value float32, selectors ...*common.Tag) Float32Toggle {
	toggle := GetFloat32Toggle(name, selectors...)
	toggle.Set(value)

	return toggle
}

func SetBoolToggle(name string, value bool, selectors ...*common.Tag) BoolToggle {
	toggle := GetBoolToggle(name, selectors...)
	toggle.Set(value)

	return toggle
}

func SetObjectToggle(name string, value map[string]any, selectors ...*common.Tag) ObjectToggle {
	toggle := GetObjectToggle(name, selectors...)
	toggle.Set(value)

	return toggle
}

func SetToggle(name, kind string, value interface{}, selectors ...*common.Tag) (Toggle, error) {
	switch kind {
	case KindString:
		val, ok := value.(string)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetStringToggle(name, val, selectors...), nil
	case KindInt:
		val, ok := value.(int)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetIntToggle(name, val, selectors...), nil
	case KindInt64:
		val, ok := value.(int64)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetInt64Toggle(name, val, selectors...), nil
	case KindFloat64:
		val, ok := value.(float64)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetFloat64Toggle(name, val, selectors...), nil
	case KindFloat32:
		val, ok := value.(float32)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetFloat32Toggle(name, val, selectors...), nil
	case KindBool:
		val, ok := value.(bool)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetBoolToggle(name, val, selectors...), nil
	case KindObject:
		val, ok := value.(map[string]interface{})

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetObjectToggle(name, val, selectors...), nil
	default:
		return nil, fmt.Errorf("unknown kind: %s", kind)
	}
}

func GetStringToggle(name string, selectors ...*common.Tag) StringToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == KindString && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(StringToggle)

		if !ok {
			t = newStringToggle(name, "", selectors)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newStringToggle(name, "", selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetIntToggle(name string, selectors ...*common.Tag) IntToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == KindInt && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(IntToggle)

		if !ok {
			t = newIntToggle(name, 0, selectors)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newIntToggle(name, 0, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetInt64Toggle(name string, selectors ...*common.Tag) Int64Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == KindInt64 && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Int64Toggle)

		if !ok {
			t = newInt64Toggle(name, 0, selectors)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newInt64Toggle(name, 0, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetFloat64Toggle(name string, selectors ...*common.Tag) Float64Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == KindFloat64 && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Float64Toggle)

		if !ok {
			t = newFloat64Toggle(name, 0, selectors)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newFloat64Toggle(name, 0, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetFloat32Toggle(name string, selectors ...*common.Tag) Float32Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == KindFloat32 && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Float32Toggle)

		if !ok {
			t = newFloat32Toggle(name, 0, selectors)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newFloat32Toggle(name, 0, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetBoolToggle(name string, selectors ...*common.Tag) BoolToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == KindBool && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(BoolToggle)

		if !ok {
			t = newBoolToggle(name, false, selectors)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newBoolToggle(name, false, selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetObjectToggle(name string, selectors ...*common.Tag) ObjectToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Type() == KindObject && toggle.Name() == name && toggle.Matches(selectors...)
	})

	if len(matches) > 0 {
		t, ok := matches[0].(ObjectToggle)

		if !ok {
			t = newObjectToggle(name, make(map[string]any), selectors)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newObjectToggle(name, make(map[string]any), selectors)
		typedToggles = append(typedToggles, t)

		return t
	}
}
