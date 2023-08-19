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
