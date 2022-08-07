package types

import (
	"fmt"
	"strconv"
)

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

func NewStringArrayValue(v *[]string) Value {
	return &stringArrayValue{
		value: v,
	}
}

type Value interface {
	Set(value interface{}) error
	Get() interface{}
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
	default:
		return fmt.Errorf("cannot convert interface to string: %T", value)
	}
	return nil
}

type stringArrayValue struct {
	value *[]string
}

func (s *stringArrayValue) Get() interface{} {
	return *s.value
}

func (s *stringArrayValue) Set(value interface{}) error {
	switch t := value.(type) {
	case []string:
		*s.value = t
	default:
		return fmt.Errorf("cannot convert interface to []string: %T", value)
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
	default:
		return fmt.Errorf("cannot convert interface to bool: %T", value)
	}
	return nil
}
