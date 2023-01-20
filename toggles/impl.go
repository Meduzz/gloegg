package toggles

import (
	"fmt"

	"github.com/Meduzz/gloegg/types"
)

type (
	featureToggle struct {
		name      string
		selectors map[string]interface{}
		value     interface{}
	}
)

func newToggle(name string, selectors map[string]interface{}, value interface{}) types.Toggle {
	return &featureToggle{
		name:      name,
		selectors: selectors,
		value:     value,
	}
}

func (f *featureToggle) Matches(selectors map[string]interface{}) bool {
	for key, value := range selectors {
		v2, ok := f.selectors[key]

		if !ok {
			return false
		}

		if value != v2 {
			return false
		}
	}

	return true
}

func (f *featureToggle) Name() string {
	return f.name
}

func (f *featureToggle) Value() interface{} {
	return f.value
}

func (f *featureToggle) SetValue(val interface{}) {
	f.value = val
}

func (f *featureToggle) GetString(def string) string {
	it, ok := f.value.(string)

	if ok {
		def = it
	}

	return def
}

func (f *featureToggle) GetInt(def int) int {
	it, ok := f.value.(int)

	if ok {
		def = it
	}

	return def
}

func (f *featureToggle) GetInt64(def int64) int64 {
	it, ok := f.value.(int64)

	if ok {
		def = it
	}

	return def
}

func (f *featureToggle) GetFloat32(def float32) float32 {
	it, ok := f.value.(float32)

	if ok {
		def = it
	}

	return def
}

func (f *featureToggle) GetFloat64(def float64) float64 {
	it, ok := f.value.(float64)

	if ok {
		def = it
	}

	return def
}

func (f *featureToggle) GetBool(def bool) bool {
	it, ok := f.value.(bool)

	if ok {
		def = it
	}

	return def
}

func (f *featureToggle) ContainsString(needle string) bool {
	array, ok := f.value.([]string)

	if !ok {
		return false
	}

	for _, it := range array {
		if it == needle {
			return true
		}
	}

	return false
}

func (f *featureToggle) ContainsInt(needle int) bool {
	array, ok := f.value.([]int)

	if !ok {
		return false
	}

	for _, it := range array {
		if it == needle {
			return true
		}
	}

	return false
}

func (f *featureToggle) ContainsInt64(needle int64) bool {
	array, ok := f.value.([]int64)

	if !ok {
		return false
	}

	for _, it := range array {
		if it == needle {
			return true
		}
	}

	return false
}

func (f *featureToggle) ContainsFloat32(needle float32) bool {
	array, ok := f.value.([]float32)

	if !ok {
		return false
	}

	for _, it := range array {
		if it == needle {
			return true
		}
	}

	return false
}

func (f *featureToggle) ContainsFloat64(needle float64) bool {
	array, ok := f.value.([]float64)

	if !ok {
		return false
	}

	for _, it := range array {
		if it == needle {
			return true
		}
	}

	return false
}

func (f *featureToggle) IsMoreInt(it int) bool {
	val, ok := f.value.(int)

	if !ok {
		return false
	}

	return val > it
}

func (f *featureToggle) IsLessInt(it int) bool {
	val, ok := f.value.(int)

	if !ok {
		fmt.Println("could not cast to int")
		return false
	}

	return val < it
}

func (f *featureToggle) IsMoreInt64(it int64) bool {
	val, ok := f.value.(int64)

	if !ok {
		return false
	}

	return val > it
}

func (f *featureToggle) IsLessInt64(it int64) bool {
	val, ok := f.value.(int64)

	if !ok {
		return false
	}

	return val < it
}

func (f *featureToggle) IsMoreFloat32(it float32) bool {
	val, ok := f.value.(float32)

	if !ok {
		return false
	}

	return val > it
}

func (f *featureToggle) IsLessFloat32(it float32) bool {
	val, ok := f.value.(float32)

	if !ok {
		return false
	}

	return val < it
}

func (f *featureToggle) IsMoreFloat64(it float64) bool {
	val, ok := f.value.(float64)

	if !ok {
		return false
	}

	return val > it
}

func (f *featureToggle) IsLessFloat64(it float64) bool {
	val, ok := f.value.(float64)

	if !ok {
		return false
	}

	return val < it
}
