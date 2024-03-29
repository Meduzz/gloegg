package common

import "github.com/Meduzz/helper/fp/slice"

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

func FindTag(tags []*Tag, key string) *Tag {
	matches := slice.Filter(tags, func(tag *Tag) bool {
		return tag.Key == key
	})

	if len(matches) == 0 {
		return nil
	}

	return matches[0]
}

func (t *Tag) String() string {
	v, ok := t.Value.(string)

	if !ok {
		return ""
	}

	return v
}

func (t *Tag) Int() int {
	v, ok := t.Value.(int)

	if !ok {
		return 0
	}

	return v
}

func (t *Tag) Int64() int64 {
	v, ok := t.Value.(int64)

	if !ok {
		return 0
	}

	return v
}

func (t *Tag) Bool() bool {
	v, ok := t.Value.(bool)

	if !ok {
		return false
	}

	return v
}
