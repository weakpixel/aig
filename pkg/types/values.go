package types

import (
	"fmt"
	"strconv"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
	"github.com/zclconf/go-cty/cty/gocty"
)

type Value interface {
	Set(value interface{}) error
	Get() interface{}
}

func NewStringValue(v *string) Value {
	return &stringValue{
		value: v,
	}
}

func NewBoolValue(v *bool) Value {
	return &boolValue{
		value: v,
	}
}

func NewIntValue(v *int) Value {
	return &intValue{
		value: v,
	}
}

func NewFloat64Value(v *float64) Value {
	return &float64Value{
		value: v,
	}
}

func NewStringListValue(v *[]string) Value {
	return &stringListValue{
		value: v,
	}
}

func NewIntListValue(v *[]int) Value {
	return &intListValue{
		value: v,
	}
}

func NewStringMapValue(v *map[string]string) Value {
	return &stringMapValue{
		value: v,
	}
}

type stringValue struct {
	value *string
}

func (s *stringValue) Get() interface{} {
	return *s.value
}

func (s *stringValue) Set(value interface{}) error {
	switch t := value.(type) {
	case string:
		*s.value = t
	case int:
		*s.value = strconv.Itoa(t)
	case bool:
		if t {
			*s.value = "true"
		} else {
			*s.value = "false"
		}
	case cty.Value:
		v, err := convert.Convert(value.(cty.Value), cty.String)
		if err != nil {
			return err
		}
		*s.value = v.AsString()

	default:
		return fmt.Errorf("cannot convert interface to string: %T", value)
	}
	return nil
}

type stringListValue struct {
	value *[]string
}

func (s *stringListValue) Get() interface{} {
	return *s.value
}

func (s *stringListValue) Set(value interface{}) error {
	switch t := value.(type) {
	case []string:
		*s.value = t
	case cty.Value:
		ty := t.Type()
		if ty.IsMapType() {
			return fmt.Errorf("cannot convert interface to []string: %T (is map) -> type: %s", value, ty.FriendlyName())
		}
		res, err := toStringList(t)
		if err != nil {
			return fmt.Errorf("cannot convert interface to []string: %T", value)
		}
		*s.value = res
	default:
		return fmt.Errorf("cannot convert interface to []string: %T", value)
	}
	return nil
}

type intListValue struct {
	value *[]int
}

func (s *intListValue) Get() interface{} {
	return *s.value
}

func (s *intListValue) Set(value interface{}) error {
	switch t := value.(type) {
	case []int:
		*s.value = t
	case cty.Value:
		ty := t.Type()
		if ty.IsMapType() {
			return fmt.Errorf("cannot convert interface to []int: %T (is map) -> type: %s", value, ty.FriendlyName())
		}
		res, err := toIntList(t)
		if err != nil {
			return fmt.Errorf("cannot convert interface to []int: %T", value)
		}
		*s.value = res
	default:
		return fmt.Errorf("cannot convert interface to []int: %T", value)
	}
	return nil
}

type boolValue struct {
	value *bool
}

func (s *boolValue) Get() interface{} {
	return *s.value
}

func (s *boolValue) Set(value interface{}) error {
	switch t := value.(type) {
	case string:
		if t == "yes" {
			*s.value = true
		} else if t == "no" {
			*s.value = false
		} else {
			b, err := strconv.ParseBool(t)
			if err != nil {
				return err
			}
			*s.value = b
		}
	case bool:
		*s.value = t
	case cty.Value:
		i, err := convert.Convert(t, cty.Bool)
		if err != nil {
			return fmt.Errorf("cannot convert value to bool: %s", err)
		}
		err = gocty.FromCtyValue(i, s.value)
		if err != nil {
			return fmt.Errorf("cannot convert value to bool: %s", err)
		}
	default:
		return fmt.Errorf("cannot convert interface to bool: %T", value)
	}
	return nil
}

type intValue struct {
	value *int
}

func (s *intValue) Get() interface{} {
	return *s.value
}

func (s *intValue) Set(value interface{}) error {
	switch t := value.(type) {
	case int:
		*s.value = t
	case string:
		b, err := strconv.Atoi(t)
		if err != nil {
			return err
		}
		*s.value = b
	case bool:
		if t {
			*s.value = 1
		} else {
			*s.value = 0
		}
	case cty.Value:
		i, err := convert.Convert(t, cty.Number)
		if err != nil {
			return fmt.Errorf("cannot convert value to int: %s", err)
		}
		err = gocty.FromCtyValue(i, s.value)
		if err != nil {
			return fmt.Errorf("cannot convert value to int: %s", err)
		}
	default:
		return fmt.Errorf("cannot convert interface to int: %T", value)
	}
	return nil
}

type float64Value struct {
	value *float64
}

func (s *float64Value) Get() interface{} {
	return *s.value
}

func (s *float64Value) Set(value interface{}) error {
	switch t := value.(type) {
	case float64:
		*s.value = t
	case string:
		f, err := strconv.ParseFloat(t, 64)
		if err != nil {
			return err
		}
		*s.value = f
	case cty.Value:
		i, err := convert.Convert(t, cty.Number)
		if err != nil {
			return fmt.Errorf("cannot convert value to float: %s", err)
		}
		err = gocty.FromCtyValue(i, s.value)
		if err != nil {
			return fmt.Errorf("cannot convert value to float: %s", err)
		}
	default:
		return fmt.Errorf("cannot convert interface to float: %T", value)
	}
	return nil
}

type stringMapValue struct {
	value *map[string]string
}

func (s *stringMapValue) Get() interface{} {
	return *s.value
}

func (s *stringMapValue) Set(value interface{}) error {
	switch t := value.(type) {
	case map[string]string:
		*s.value = t
	case cty.Value:
		re, err := toStringMap(t)
		if err != nil {
			return fmt.Errorf("cannot convert value to map[string]string: %s", err)
		}
		for k := range *s.value {
			delete(*s.value, k)
		}
		*s.value = re
	default:
		return fmt.Errorf("cannot convert interface to map[string]string: %T", value)
	}
	return nil
}
