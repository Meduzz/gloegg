package toggles

import "github.com/Meduzz/gloegg/types"

var (
	featureToggles       []types.Toggle
	toggleUpdatedHandler func(types.Toggle)
)

// SetToggle lets us create or update a feature toggle and its value
func SetToggle(name string, value interface{}, selectors ...*types.Tag) types.Toggle {
	exists := false
	var t types.Toggle

	for _, it := range featureToggles {
		if it.Name() == name && it.Matches(types.AsMap(selectors...)) {
			exists = true
			it.SetValue(value)
			t = it

			if toggleUpdatedHandler != nil {
				toggleUpdatedHandler(it)
			}
		}
	}

	if !exists {
		t = newToggle(name, types.AsMap(selectors...), value)
		featureToggles = append(featureToggles, t)
	}

	return t
}

// GetToggle lets us fetch a feature toggle, note that both name and selectors must match!
func GetToggle(name string, selectors ...*types.Tag) types.Toggle {
	for _, it := range featureToggles {
		if it.Name() == name && it.Matches(types.AsMap(selectors...)) {
			return it
		}
	}

	return nil
}

// SetToggleUpdatedHandler allows us to set a listener to when toggles are updated
func SetToggleUpdatedHandler(handler func(types.Toggle)) {
	toggleUpdatedHandler = handler
}
