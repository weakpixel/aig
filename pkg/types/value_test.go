package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

}

func TestStringArrayValue(t *testing.T) {
	val := []string{}
	vv := NewStringArrayValue(&val)
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
}
