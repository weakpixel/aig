package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestStringValue(t *testing.T) {
	val := ""
	vv := NewStringValue(&val)

	err := vv.Set("newValue")
	if assert.NoError(t, err) {
		assert.Equal(t, "newValue", val)
		assert.Equal(t, "newValue", vv.Get())
	}

	err = vv.Set(true)
	if assert.NoError(t, err) {
		assert.Equal(t, "true", val)
		assert.Equal(t, "true", vv.Get())
	}

	err = vv.Set(false)
	if assert.NoError(t, err) {
		assert.Equal(t, "false", val)
		assert.Equal(t, "false", vv.Get())
	}

	err = vv.Set(10)
	if assert.NoError(t, err) {
		assert.Equal(t, "10", val)
		assert.Equal(t, "10", vv.Get())
	}

	err = vv.Set(cty.StringVal("cty-value"))
	if assert.NoError(t, err) {
		assert.Equal(t, "cty-value", val)
		assert.Equal(t, "cty-value", vv.Get())
	}

}

func TestStringArrayValue(t *testing.T) {
	val := []string{}
	vv := NewStringListValue(&val)
	newVal := []string{
		"a",
		"b",
	}
	err := vv.Set(newVal)
	if assert.NoError(t, err) {
		assert.Equal(t, newVal, val)
		assert.Equal(t, newVal, vv.Get())
	}

	tuple := cty.TupleVal([]cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
	})

	err = vv.Set(tuple)
	if assert.NoError(t, err) {
		assert.Equal(t, newVal, val)
		assert.Equal(t, newVal, vv.Get())
	}

	list := cty.ListVal([]cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
	})
	err = vv.Set(list)
	if assert.NoError(t, err) {
		assert.Equal(t, newVal, val)
		assert.Equal(t, newVal, vv.Get())
	}

}

func TestIntArrayValue(t *testing.T) {
	val := []int{}
	vv := NewIntListValue(&val)
	newVal := []int{
		1,
		2,
	}
	err := vv.Set(newVal)
	if assert.NoError(t, err) {
		assert.Equal(t, newVal, val)
		assert.Equal(t, newVal, vv.Get())
	}

	tuple := cty.TupleVal([]cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(2),
	})

	err = vv.Set(tuple)
	if assert.NoError(t, err) {
		assert.Equal(t, newVal, val)
		assert.Equal(t, newVal, vv.Get())
	}

	list := cty.ListVal([]cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(2),
	})
	err = vv.Set(list)
	if assert.NoError(t, err) {
		assert.Equal(t, newVal, val)
		assert.Equal(t, newVal, vv.Get())
	}

}

func TestBoolValue(t *testing.T) {
	val := false
	vv := NewBoolValue(&val)
	err := vv.Set(true)
	if assert.NoError(t, err) {
		assert.Equal(t, true, val)
	}

	err = vv.Set(false)
	if assert.NoError(t, err) {
		assert.Equal(t, false, val)
	}

	err = vv.Set("true")
	if assert.NoError(t, err) {
		assert.Equal(t, true, val)
	}

	err = vv.Set("false")
	if assert.NoError(t, err) {
		assert.Equal(t, false, val)
	}

	err = vv.Set(cty.BoolVal(true))
	if assert.NoError(t, err) {
		assert.Equal(t, true, val)
	}

}

func TestIntValue(t *testing.T) {
	val := 0
	vv := NewIntValue(&val)
	err := vv.Set(true)
	if assert.NoError(t, err) {
		assert.Equal(t, 1, val)
	}

	err = vv.Set(false)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, val)
	}

	err = vv.Set("10")
	if assert.NoError(t, err) {
		assert.Equal(t, 10, val)
	}

	err = vv.Set(40)
	if assert.NoError(t, err) {
		assert.Equal(t, 40, val)
	}

	err = vv.Set(cty.NumberIntVal(66))
	if assert.NoError(t, err) {
		assert.Equal(t, 66, val)
	}

}

func TestFloat64Value(t *testing.T) {
	val := 0.0
	vv := NewFloat64Value(&val)
	err := vv.Set("101.23")
	if assert.NoError(t, err) {
		assert.Equal(t, 101.23, val)
	}

	err = vv.Set(40.345)
	if assert.NoError(t, err) {
		assert.Equal(t, 40.345, val)
	}

	err = vv.Set(cty.NumberFloatVal(111.345))
	if assert.NoError(t, err) {
		assert.Equal(t, 111.345, val)
	}

}

func TestStringMapValue(t *testing.T) {
	val := map[string]string{}
	vv := NewStringMapValue(&val)
	newVal := map[string]string{
		"a": "a value",
		"b": "b value",
	}

	err := vv.Set(newVal)
	if assert.NoError(t, err) {
		assert.Equal(t, newVal, val)
		assert.Equal(t, newVal, vv.Get())
	}

	m := cty.MapVal(map[string]cty.Value{
		"c": cty.StringVal("c value"),
		"d": cty.StringVal("d value"),
	})
	newVal = map[string]string{
		"c": "c value",
		"d": "d value",
	}

	err = vv.Set(m)
	if assert.NoError(t, err) {
		assert.Equal(t, newVal, val)
		assert.Equal(t, newVal, vv.Get())
		fmt.Println(vv.Get())
	}

}
