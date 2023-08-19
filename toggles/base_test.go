package toggles

import (
	"testing"

	"github.com/Meduzz/gloegg/common"
)

func TestTypes(t *testing.T) {
	s := SetStringToggle("string", "asdf")
	sVal := s.DefaultValue("fail")

	if sVal != "asdf" {
		t.Errorf("strVal was not 'asdf' but '%s'", sVal)
		t.FailNow()
	}

	s.Set("test")
	sVal = s.DefaultValue("fail")

	if sVal != "test" {
		t.Errorf("strVal was not 'test' but '%s'", sVal)
		t.FailNow()
	}

	i := SetIntToggle("int", 1)
	iVal := i.DefaultValue(0)

	if iVal != 1 {
		t.Errorf("iVal was not '1' but '%d'", iVal)
		t.FailNow()
	}

	if !i.MoreThan(0) {
		t.Errorf("0 was not less than iVal(%d)", iVal)
		t.FailNow()
	}

	if !i.LessThan(2) {
		t.Errorf("2 was not more than iVal(%d)", iVal)
		t.FailNow()
	}

	i.Set(2)

	iVal = i.DefaultValue(0)

	if iVal != 2 {
		t.Errorf("iVal was not '2' but '%d'", iVal)
		t.FailNow()
	}

	i64 := SetInt64Toggle("int64", int64(1))
	i64Val := i64.DefaultValue(0)

	if i64Val != int64(1) {
		t.Errorf("i64Val was not '1' but '%d'", i64Val)
		t.FailNow()
	}

	if !i64.MoreThan(0) {
		t.Errorf("0 was not less that i64Val(%d)", i64Val)
		t.FailNow()
	}

	if !i64.LessThan(2) {
		t.Errorf("2 was not more than i64Val(%d)", i64Val)
		t.FailNow()
	}

	i64.Set(2)
	i64Val = i64.DefaultValue(0)

	if i64Val != int64(2) {
		t.Errorf("i64Val was not '2' but '%d'", i64Val)
		t.FailNow()
	}

	f := SetFloat32Toggle("float", float32(0.1))
	fVal := f.DefaultValue(0.2)

	if fVal != 0.1 {
		t.Errorf("fVal was not '0.1' but '%f'", fVal)
		t.FailNow()
	}

	if !f.MoreThan(0.0) {
		t.Errorf("0.0 was not less that fVal(%f)", fVal)
		t.FailNow()
	}

	if !f.LessThan(0.2) {
		t.Errorf("0.2 was not more than fVal(%f)", fVal)
		t.FailNow()
	}

	f.Set(0.2)
	fVal = f.DefaultValue(0.0)

	if fVal != 0.2 {
		t.Errorf("fVal was not '0.2' but '%f'", fVal)
		t.FailNow()
	}

	f64 := SetFloat64Toggle("float64", float64(0.1))
	f64Val := f64.DefaultValue(0.2)

	if f64Val != float64(0.1) {
		t.Errorf("f64Val was not '0.1' but '%f'", f64Val)
		t.FailNow()
	}

	if !f64.MoreThan(0.0) {
		t.Errorf("0.0 was not less that f64Val(%f)", f64Val)
		t.FailNow()
	}

	if !f64.LessThan(0.2) {
		t.Errorf("0.2 was not more than f64Val(%f)", f64Val)
		t.FailNow()
	}

	f64.Set(0.2)
	f64Val = f64.DefaultValue(0.0)

	if f64Val != float64(0.2) {
		t.Errorf("f64Val was not '0.2' but '%f'", f64Val)
		t.FailNow()
	}

	b := SetBoolToggle("bool", true)
	bVal := b.Value()

	if !bVal {
		t.Error("bVal was not true but false")
		t.FailNow()
	}

	b.Set(false)
	bVal = b.Value()

	if bVal {
		t.Error("bVal was not false but true")
		t.FailNow()
	}

	data := make(map[string]any)
	o := SetObjectToggle("object", data)

	oVal := o.Value()

	if len(oVal) != len(data) {
		t.Error("oVal was not equal to data")
		t.FailNow()
	}

	o.SetField("s", "test")
	o.SetField("i", 2)
	o.SetField("i64", int64(2))
	o.SetField("f", float32(0.2))
	o.SetField("f64", float64(0.2))
	o.SetField("b", false)

	sVal = o.DefaultString("s", "")

	if sVal != "test" {
		t.Errorf("strVal was not 'test' but '%s'", sVal)
		t.FailNow()
	}

	iVal = o.DefaultInt("i", 0)

	if iVal != 2 {
		t.Errorf("iVal was not '2' but '%d'", iVal)
		t.FailNow()
	}

	i64Val = o.DefaultInt64("i64", 0)

	if i64Val != int64(2) {
		t.Errorf("i64Val was not '2' but '%d'", i64Val)
		t.FailNow()
	}

	fVal = o.DefaultFloat("f", float32(0.0))

	if fVal != float32(0.2) {
		t.Errorf("fVal was not '0.2' but '%f'", fVal)
		t.FailNow()
	}

	f64Val = o.DefaultFloat64("f64", float64(0.0))

	if f64Val != float64(0.2) {
		t.Errorf("f64Val was not '0.2' but '%f'", f64Val)
		t.FailNow()
	}

	bVal = o.DefaultBool("b", true)

	if bVal {
		t.Error("bVal was not false but true")
		t.FailNow()
	}
}

func TestSelectors(t *testing.T) {
	subject := SetStringToggle("name", "good", common.Pair("number", 8), common.Pair("strings", "cool"))

	// order should not matter
	if !subject.Matches(common.Pair("strings", "cool"), common.Pair("number", 8)) {
		t.Error("selectors did not match as expected")
	}

	// only key match
	if subject.Matches(common.Pair("strings", "cool!")) {
		t.Error("unexpected match")
	}

	// only "partial tag" match
	if subject.Matches(common.Pair("strings", "cool!"), common.Pair("number", 8)) {
		t.Error("unexpected match")
	}

	// full "partial tag" match
	if !subject.Matches(common.Pair("number", 8)) {
		t.Error("expected tags to match")
	}

}

func TestUpdateValue(t *testing.T) {
	subject := SetStringToggle("value", "good")
	subject.Set("bad")

	val := subject.DefaultValue("test")

	if val != "bad" {
		t.Errorf("val was not 'bad' but '%s'", val)
		t.FailNow()
	}
}
