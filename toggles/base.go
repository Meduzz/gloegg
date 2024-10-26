package toggles

import (
	"fmt"

	"github.com/Meduzz/helper/fp/slice"
)

var (
	typedToggles []Toggle
	callbacks    chan *UpdatedToggle
	subscribers  []Subscriber
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

func init() {
	callbacks = make(chan *UpdatedToggle, 100)
	subscribers = make([]Subscriber, 0)

	go func() {
		for update := range callbacks {
			slice.ForEach(subscribers, func(sub Subscriber) {
				sub(update)
			})
		}
	}()
}

func Subscribe(handler Subscriber) {
	subscribers = append(subscribers, handler)
}

func SetStringToggle(name, value string) StringToggle {
	toggle := GetStringToggle(name)
	toggle.Set(value)

	return toggle
}

func SetIntToggle(name string, value int) IntToggle {
	toggle := GetIntToggle(name)
	toggle.Set(value)

	return toggle
}

func SetInt64Toggle(name string, value int64) Int64Toggle {
	toggle := GetInt64Toggle(name)
	toggle.Set(value)

	return toggle
}

func SetFloat64Toggle(name string, value float64) Float64Toggle {
	toggle := GetFloat64Toggle(name)
	toggle.Set(value)

	return toggle
}

func SetFloat32Toggle(name string, value float32) Float32Toggle {
	toggle := GetFloat32Toggle(name)
	toggle.Set(value)

	return toggle
}

func SetBoolToggle(name string, value bool) BoolToggle {
	toggle := GetBoolToggle(name)
	toggle.Set(value)

	return toggle
}

func SetObjectToggle(name string, value map[string]any) ObjectToggle {
	toggle := GetObjectToggle(name)
	toggle.Set(value)

	return toggle
}

func SetToggle(name, kind string, value interface{}) (Toggle, error) {
	switch kind {
	case KindString:
		val, ok := value.(string)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetStringToggle(name, val), nil
	case KindInt:
		val, ok := value.(int)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetIntToggle(name, val), nil
	case KindInt64:
		val, ok := value.(int64)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetInt64Toggle(name, val), nil
	case KindFloat64:
		val, ok := value.(float64)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetFloat64Toggle(name, val), nil
	case KindFloat32:
		val, ok := value.(float32)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetFloat32Toggle(name, val), nil
	case KindBool:
		val, ok := value.(bool)

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetBoolToggle(name, val), nil
	case KindObject:
		val, ok := value.(map[string]interface{})

		if !ok {
			return nil, fmt.Errorf("could not convert value to type %s", kind)
		}

		return SetObjectToggle(name, val), nil
	default:
		return nil, fmt.Errorf("unknown kind: %s", kind)
	}
}

func GetStringToggle(name string) StringToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Kind() == KindString && toggle.Name() == name
	})

	if len(matches) > 0 {
		t, ok := matches[0].(StringToggle)

		if !ok {
			t = newStringToggle(name, "", callbacks)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newStringToggle(name, "", callbacks)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetIntToggle(name string) IntToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Kind() == KindInt && toggle.Name() == name
	})

	if len(matches) > 0 {
		t, ok := matches[0].(IntToggle)

		if !ok {
			t = newIntToggle(name, 0, callbacks)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newIntToggle(name, 0, callbacks)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetInt64Toggle(name string) Int64Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Kind() == KindInt64 && toggle.Name() == name
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Int64Toggle)

		if !ok {
			t = newInt64Toggle(name, 0, callbacks)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newInt64Toggle(name, 0, callbacks)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetFloat64Toggle(name string) Float64Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Kind() == KindFloat64 && toggle.Name() == name
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Float64Toggle)

		if !ok {
			t = newFloat64Toggle(name, 0, callbacks)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newFloat64Toggle(name, 0, callbacks)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetFloat32Toggle(name string) Float32Toggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Kind() == KindFloat32 && toggle.Name() == name
	})

	if len(matches) > 0 {
		t, ok := matches[0].(Float32Toggle)

		if !ok {
			t = newFloat32Toggle(name, 0, callbacks)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newFloat32Toggle(name, 0, callbacks)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetBoolToggle(name string) BoolToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Kind() == KindBool && toggle.Name() == name
	})

	if len(matches) > 0 {
		t, ok := matches[0].(BoolToggle)

		if !ok {
			t = newBoolToggle(name, false, callbacks)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newBoolToggle(name, false, callbacks)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func GetObjectToggle(name string) ObjectToggle {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Kind() == KindObject && toggle.Name() == name
	})

	if len(matches) > 0 {
		t, ok := matches[0].(ObjectToggle)

		if !ok {
			t = newObjectToggle(name, make(map[string]any), callbacks)
			typedToggles = append(typedToggles, t)
		}

		return t
	} else {
		t := newObjectToggle(name, make(map[string]any), callbacks)
		typedToggles = append(typedToggles, t)

		return t
	}
}

func RemoveToggle(name, kind string) {
	matches := slice.Filter(typedToggles, func(toggle Toggle) bool {
		return toggle.Kind() == kind && toggle.Name() == name
	})

	pool := typedToggles
	keepers := make([]Toggle, 0)

	slice.ForEach(matches, func(it Toggle) {
		for _, t := range pool {
			if t.Name() == it.Name() {
				if t.Kind() == it.Kind() {
					continue
				}
			}

			keepers = append(keepers, t)
		}

		pool = keepers
	})

	typedToggles = pool
}
