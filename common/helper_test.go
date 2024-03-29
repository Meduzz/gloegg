package common

import (
	"testing"
)

func TestFindTags(t *testing.T) {
	subjects := make([]*Tag, 0)
	subjects = append(subjects, Pair("key1", "value1"), Pair("key2", "value2"))

	match := FindTag(subjects, "key1")

	if match == nil {
		t.Errorf("Tag with key=key1 was not found.")
		t.Fail()
	}

	match = FindTag(subjects, "asdf")

	if match != nil {
		t.Errorf("A Tag with key=asdf was found.")
		t.Fail()
	}
}

func TestPairs(t *testing.T) {
	subject := Pairs("first", 1, "second", 2, "ignored")

	if len(subject) != 2 {
		t.Errorf("invalid length of subject, expected 2 have %d", len(subject))
	}

	first := FindTag(subject, "first")

	if first.Int() != 1 {
		t.Errorf("first value was not 1 but %d", first.Int())
	}

	second := FindTag(subject, "second")

	if second.Int() != 2 {
		t.Errorf("second value was not 2 but %d", second.Int())
	}
}
