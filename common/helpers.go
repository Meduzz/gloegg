package common

import (
	"log/slog"

	"github.com/Meduzz/helper/fp/slice"
)

// Pair create a tag
func Pair(key string, value interface{}) *Tag {
	return &Tag{key, value}
}

func Pairs(attrs ...any) []*Tag {
	arrayOfPairs := slice.Partition(attrs, 2)

	return slice.Fold(arrayOfPairs, make([]*Tag, 0), func(pair []any, agg []*Tag) []*Tag {
		key, ok := pair[0].(string)

		if ok && len(pair) == 2 {
			return append(agg, Pair(key, pair[1]))
		}

		return agg
	})
}

// AsMap turn a bunch of tags into map[string]interface{}
func AsMap(tags ...*Tag) map[string]interface{} {
	it := make(map[string]interface{})

	for _, tag := range tags {
		it[tag.Key] = tag.Value
	}

	return it
}

// AsTags turns a map into a list of tags.
func AsTags(in map[string]interface{}) []*Tag {
	out := make([]*Tag, 0)

	for k, v := range in {
		out = append(out, Pair(k, v))
	}

	return out
}

func (t Tags) ToSlog() []slog.Attr {
	return slice.Map(t, func(tag *Tag) slog.Attr {
		return slog.Any(tag.Key, tag.Value)
	})
}
