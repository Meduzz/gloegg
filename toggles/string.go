package toggles

import (
	"strings"

	"github.com/Meduzz/gloegg/common"
)

type (
	stringToggle struct {
		name     string
		value    string
		metadata []*common.Tag
	}
)

func newStringToggle(name, value string, metadata []*common.Tag) StringToggle {
	return &stringToggle{name, value, metadata}
}

func (s *stringToggle) Name() string {
	return s.name
}

func (s *stringToggle) Matches(needle ...*common.Tag) bool {
	return matches(needle, s.metadata)
}

func (s *stringToggle) Type() string {
	return KindString
}

func (s *stringToggle) Value() string {
	return s.value
}

func (s *stringToggle) Set(value string) {
	s.value = value
}

func (s *stringToggle) DefaultValue(defaultValue string) string {
	if s.value == "" {
		return defaultValue
	}

	return s.value
}

func (s *stringToggle) Equals(value string) bool {
	return s.value == value
}

func (s *stringToggle) Contains(value string) bool {
	return strings.Contains(s.value, value)
}
