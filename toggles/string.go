package toggles

import (
	"strings"
)

type (
	stringToggle struct {
		*base
	}
)

func newStringToggle(name, value string, callbacks chan *UpdatedToggle) StringToggle {
	return &stringToggle{&base{name, KindString, value, callbacks}}
}

func (s *stringToggle) Kind() string {
	return KindString
}

func (s *stringToggle) Value() string {
	result, ok := s.value.(string)

	if !ok {
		return ""
	}

	return result
}

func (s *stringToggle) Set(value string) {
	s.value = value
	s.callback <- s.Updated()
}

func (s *stringToggle) DefaultValue(defaultValue string) string {
	result := s.Value()

	if result == "" {
		return defaultValue
	}

	return result
}

func (s *stringToggle) Equals(value string) bool {
	return s.value == value
}

func (s *stringToggle) Contains(value string) bool {
	return strings.Contains(s.Value(), value)
}
