package types

import (
	"context"
)

type (
	Logging interface {
		// Info - log a msg at info level, providing additional context via tags
		Info(msg string, tags ...*Tag)

		// Debug - log a msg at debug level, providing additional context via tags
		Debug(msg string, tags ...*Tag)

		// Warn - log a msg at warn level, providing additional context via tags
		Warn(msg string, tags ...*Tag)

		// Error - log a msg at error level, providing additional context via tags
		Error(msg string, err error, tags ...*Tag)

		// TODO InDebug(run func(logger)) style methods? For "expensive" logs.

		// Trace - create a new named trace, providing additional context via tags
		Trace(name string, tags ...*Tag) (context.Context, func(error))

		// Trace - create a new named trace, providing additional context via tags
		TraceContext(name string, parent context.Context, tags ...*Tag) (context.Context, func(error))

		// FeatureToggle - load a feature toggle
		FeatureToggle(name string, tags ...*Tag) Toggle
	}

	Trace interface {
		Start() (context.Context, func(error))
		ID() string
		Name() string
	}

	Toggle interface {
		Matches(map[string]interface{}) bool
		Name() string
		Value() interface{}
		SetValue(interface{})

		// fetch the actual value
		GetString(string) string
		GetInt(int) int
		GetInt64(int64) int64
		GetFloat32(float32) float32
		GetFloat64(float64) float64
		GetBool(bool) bool

		// do comparasion against the value(s)
		// multi values
		ContainsString(string) bool
		ContainsInt(int) bool
		ContainsInt64(int64) bool
		ContainsFloat32(float32) bool
		ContainsFloat64(float64) bool

		// single value
		IsMoreInt(int) bool
		IsLessInt(int) bool
		IsMoreInt64(int64) bool
		IsLessInt64(int64) bool
		IsMoreFloat32(float32) bool
		IsLessFloat32(float32) bool
		IsMoreFloat64(float64) bool
		IsLessFloat64(float64) bool
	}

	Tag struct {
		Key   string
		Value interface{}
	}

	Stack struct {
		File string `json:"file"`
		Line int    `json:"line"`
		Func string `json:"function"`
	}
)

// Pair create a tag
func Pair(key string, value interface{}) *Tag {
	return &Tag{key, value}
}

// AsMap turn a bunch of tags into map[string]interface{}
func AsMap(tags ...*Tag) map[string]interface{} {
	it := make(map[string]interface{})

	for _, tag := range tags {
		it[tag.Key] = tag.Value
	}

	return it
}
