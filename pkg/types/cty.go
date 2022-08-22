package types

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
	"github.com/zclconf/go-cty/cty/gocty"
)

type CtyValue interface {
	Type() cty.Type
}

func toStringList(v cty.Value) ([]string, error) {
	newval, err := convert.Convert(v, cty.List(cty.String))
	if err != nil {
		return nil, fmt.Errorf("cannot convert tuple value: %s", err)
	}
	itr := newval.ElementIterator()
	list := []string{}
	for itr.Next() {
		_, e := itr.Element()
		e, err := convert.Convert(e, cty.String)
		if err != nil {
			return nil, fmt.Errorf("cannot convert tuple value: %s", err)
		}
		list = append(list, e.AsString())
	}
	return list, nil
}

func toIntList(v cty.Value) ([]int, error) {
	newval, err := convert.Convert(v, cty.List(cty.Number))
	if err != nil {
		return nil, fmt.Errorf("cannot convert value to list: %s", err)
	}
	itr := newval.ElementIterator()
	list := []int{}
	for itr.Next() {
		_, e := itr.Element()
		e, err := convert.Convert(e, cty.Number)
		if err != nil {
			return nil, fmt.Errorf("cannot convert value to int: %s", err)
		}
		var i int
		err = gocty.FromCtyValue(e, &i)
		if err != nil {
			return nil, fmt.Errorf("cannot convert value to int: %s", err)
		}
		list = append(list, i)
	}
	return list, nil
}

func toStringMap(v cty.Value) (map[string]string, error) {
	result := map[string]string{}
	if v.Type().IsMapType() {
		for k, v := range v.AsValueMap() {

			v, err := convert.Convert(v, cty.String)
			if err != nil {
				return nil, fmt.Errorf("cannot convert value value: %s", err)
			}
			result[k] = v.AsString()
		}
	}
	return result, nil
}
