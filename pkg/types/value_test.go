package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValue(t *testing.T) {
	val := ""
	vv := &stringValue{
		value: &val,
	}

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

}

func TestStringArrayValue(t *testing.T) {
	val := []string{}
	vv := &stringArrayValue{
		value: &val,
	}
	newVal := []string{
		"a",
		"b",
	}
	err := vv.Set(newVal)
	if assert.NoError(t, err) {
		assert.Equal(t, newVal, val)
		assert.Equal(t, newVal, vv.Get())
	}

}

func TestBoolValue(t *testing.T) {
	val := false
	vv := &boolValue{
		value: &val,
	}
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
}

func TestIntValue(t *testing.T) {
	val := 0
	vv := &intValue{
		value: &val,
	}
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
}
