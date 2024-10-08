package toggles

import (
	"testing"
)

func TestToggles(t *testing.T) {
	t.Run("Test Types", func(t *testing.T) {
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
	})

	t.Run("Test update value", func(t *testing.T) {
		subject := SetStringToggle("value", "good")
		subject.Set("bad")

		val := subject.DefaultValue("test")

		if val != "bad" {
			t.Errorf("val was not 'bad' but '%s'", val)
			t.FailNow()
		}
	})

	t.Run("Test set string toggle", func(t *testing.T) {
		toggle, err := SetToggle("string", KindString, "test")

		if err != nil {
			t.Errorf("creating toggle threw error: %v", err)
		}

		stringToggle, ok := toggle.(StringToggle)

		if !ok {
			t.Error("toggle was not of string kind")
		}

		if !stringToggle.Equals("test") {
			t.Errorf("toggle value was not 'test' but '%s'", stringToggle.Value())
		}

		stringToggle = GetStringToggle("string")

		if stringToggle.DefaultValue("ERROR") != "test" {
			t.Errorf("string toggle value was not 'test' but '%s'", stringToggle.Value())
		}

		toggle, err = SetToggle("string", KindString, "real")

		if err != nil {
			t.Errorf("creating toggle threw error: %v", err)
		}

		stringToggle, ok = toggle.(StringToggle)

		if !ok {
			t.Error("toggle was not of string kind")
		}

		if !stringToggle.Equals("real") {
			t.Errorf("toggle value was not 'real' but '%s'", stringToggle.Value())
		}

		stringToggle = GetStringToggle("string")

		if stringToggle.DefaultValue("ERROR") != "real" {
			t.Errorf("string toggle value was not 'real' but '%s'", stringToggle.Value())
		}
	})

	t.Run("Test set object toggle", func(t *testing.T) {
		data := make(map[string]interface{})
		data["value"] = "test"

		toggle, err := SetToggle("string", KindObject, data)

		if err != nil {
			t.Errorf("creating toggle threw error: %v", err)
		}

		objectToggle, ok := toggle.(ObjectToggle)

		if !ok {
			t.Error("toggle was not of string kind")
		}

		if objectToggle.DefaultString("value", "ERROR") != "test" {
			t.Errorf("toggle value was not 'test' but '%s'", objectToggle.Value())
		}

		objectToggle = GetObjectToggle("string")

		if objectToggle.DefaultString("value", "ERROR") != "test" {
			t.Errorf("string toggle value was not 'test' but '%s'", objectToggle.Value())
		}

		data["value"] = "real"
		toggle, err = SetToggle("string", KindObject, data)

		if err != nil {
			t.Errorf("creating toggle threw error: %v", err)
		}

		objectToggle, ok = toggle.(ObjectToggle)

		if !ok {
			t.Error("toggle was not of string kind")
		}

		if objectToggle.DefaultString("value", "ERROR") != "real" {
			t.Errorf("toggle value was not 'real' but '%s'", objectToggle.Value())
		}

		objectToggle = GetObjectToggle("string")

		if objectToggle.DefaultString("value", "ERROR") != "real" {
			t.Errorf("string toggle value was not 'real' but '%s'", objectToggle.Value())
		}
	})

	t.Run("Test remove toggle", func(t *testing.T) {
		SetStringToggle("value", "good")

		subject := GetStringToggle("value")

		if subject.Value() != "good" {
			t.Errorf("toggle value was not 'good' but '%s'", subject.Value())
		}

		RemoveToggle("value", KindString)

		subject = GetStringToggle("value")

		if subject.DefaultValue("removed") != "removed" {
			t.Errorf("toggle value was not 'removed' but '%s'", subject.Value())
		}
	})

	t.Run("Test subscribe", func(t *testing.T) {
		subject := SetStringToggle("value", "good")
		feedback := make(chan string, 1)

		Subscribe(func(ut *UpdatedToggle) {
			if ut.Kind != KindString && ut.Name != "value" {
				return
			}

			feedback <- ut.Value.(string)
		})

		subject.Set("bad")

		<-feedback // good
		result := <-feedback

		if result != "bad" {
			t.Errorf("value was not bad but %s", result)
		}

	})
}
