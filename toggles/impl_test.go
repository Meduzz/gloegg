package toggles

import (
	"testing"

	"github.com/Meduzz/gloegg/types"
)

func TestTypes(t *testing.T) {
	str := newToggle("string", nil, "asdf")
	strVal := str.GetString("fail")

	if strVal != "asdf" {
		t.Errorf("strVal was not 'asdf' but '%s'", strVal)
		t.FailNow()
	}

	i := newToggle("int", nil, 1)
	iVal := i.GetInt(0)

	if iVal != 1 {
		t.Errorf("iVal was not '1' but '%d'", iVal)
		t.FailNow()
	}

	if !i.IsMoreInt(0) {
		t.Errorf("0 was not less than iVal(%d)", iVal)
		t.FailNow()
	}

	if !i.IsLessInt(2) {
		t.Errorf("2 was not more than iVal(%d)", iVal)
		t.FailNow()
	}

	i64 := newToggle("int64", nil, int64(1))
	i64Val := i64.GetInt64(0)

	if i64Val != int64(1) {
		t.Errorf("i64Val was not '1' but '%d'", i64Val)
		t.FailNow()
	}

	if !i64.IsMoreInt64(0) {
		t.Errorf("0 was not less that i64Val(%d)", i64Val)
		t.FailNow()
	}

	if !i64.IsLessInt64(2) {
		t.Errorf("2 was not more than i64Val(%d)", i64Val)
		t.FailNow()
	}

	f := newToggle("float", nil, float32(0.1))
	fVal := f.GetFloat32(0.2)

	if fVal != 0.1 {
		t.Errorf("fVal was not '0.1' but '%f'", fVal)
		t.FailNow()
	}

	if !f.IsMoreFloat32(0.0) {
		t.Errorf("0.0 was not less that fVal(%f)", fVal)
		t.FailNow()
	}

	if !f.IsLessFloat32(0.2) {
		t.Errorf("0.2 was not more than fVal(%f)", fVal)
		t.FailNow()
	}

	f64 := newToggle("float64", nil, float64(0.1))
	f64Val := f64.GetFloat64(0.2)

	if f64Val != float64(0.1) {
		t.Errorf("f64Val was not '0.1' but '%f'", f64Val)
		t.FailNow()
	}

	if !f64.IsMoreFloat64(0.0) {
		t.Errorf("0.0 was not less that f64Val(%f)", f64Val)
		t.FailNow()
	}

	if !f64.IsLessFloat64(0.2) {
		t.Errorf("0.2 was not more than f64Val(%f)", f64Val)
		t.FailNow()
	}

	b := newToggle("bool", nil, true)
	bVal := b.GetBool(false)

	if !bVal {
		t.Error("bVal was not true but false")
		t.FailNow()
	}

	strArray := newToggle("strArray", nil, []string{"a", "b", "c"})

	if !strArray.ContainsString("a") {
		t.Error("strArray did not contain 'a'")
		t.FailNow()
	}

	if strArray.ContainsString("d") {
		t.Error("strArray did contain 'd'")
		t.FailNow()
	}

	iArray := newToggle("iArray", nil, []int{1, 2, 3})

	if !iArray.ContainsInt(1) {
		t.Error("iArray did not contain 1")
		t.FailNow()
	}

	if iArray.ContainsInt(4) {
		t.Error("iArray did contain 4")
		t.FailNow()
	}

	i64Array := newToggle("i64Array", nil, []int64{1, 2, 3})

	if !i64Array.ContainsInt64(1) {
		t.Error("i64Array did not contain 1")
		t.FailNow()
	}

	if i64Array.ContainsInt64(4) {
		t.Error("i64Array did contain 4")
		t.FailNow()
	}

	fArray := newToggle("fArray", nil, []float32{0.1, 0.2, 0.3})

	if !fArray.ContainsFloat32(0.1) {
		t.Error("fArray did not contain 0.1")
		t.FailNow()
	}

	if fArray.ContainsFloat32(0.4) {
		t.Error("fArray did contain 0.4")
		t.FailNow()
	}

	f64Array := newToggle("f64Array", nil, []float64{0.1, 0.2, 0.3})

	if !f64Array.ContainsFloat64(0.1) {
		t.Error("f64Array did not contain 0.1")
		t.FailNow()
	}

	if f64Array.ContainsFloat64(0.4) {
		t.Error("f64Array did contain 0.4")
		t.FailNow()
	}
}

func TestSelectors(t *testing.T) {
	subject := newToggle("name", types.AsMap(types.Pair("number", 8), types.Pair("strings", "cool")), "good")

	// order should not matter
	if !subject.Matches(types.AsMap(types.Pair("strings", "cool"), types.Pair("number", 8))) {
		t.Error("selectors did not match as expected")
		t.FailNow()
	}
}

func TestUpdateValue(t *testing.T) {
	subject := newToggle("value", nil, "good")
	subject.SetValue("bad")

	val := subject.GetString("test")

	if val != "bad" {
		t.Errorf("val was not 'bad' but '%s'", val)
		t.FailNow()
	}
}
