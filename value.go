package lang

import (
	"errors"
	"fmt"
)

type ValueType int

var (
	NotValidType = errors.New("invalid type")
)

const (
	ValueFunction ValueType = iota
	ValueInt
	ValueString
	ValueBool
	ValueComplex
	ValueIndirect
)

type Value struct {
	typ  ValueType
	v    interface{}
	vari variable
}

func NewValue(raw interface{}) (*Value, error) {
	switch raw.(type) {
	case int:
		return &Value{
			typ: ValueInt,
			v:   raw,
		}, nil
	case string:
		return &Value{
			typ: ValueString,
			v:   raw,
		}, nil
	case variable:
		vari := raw.(variable)
		return &Value{
			typ:  ValueIndirect,
			vari: vari,
		}, nil
	}

	return nil, errors.New("Not implemented")
}

func (v *Value) V() (interface{}, error) {
	if v.v != nil {
		return v.v, nil
	}

	value := v.vari.scope.Get(v.vari.name)
	if value == nil {
		return nil, errors.New("unknown variable")
	}

	return value.V()
}

func (v *Value) ToInt() (int, error) {
	raw, err := v.V()
	if err != nil {
		return 0, err
	}

	i, ok := raw.(int)
	if !ok {
		return 0, NotValidType
	}

	return i, nil
}

func (v *Value) MustInt() int {
	i, _ := v.ToInt()

	return i
}

func (v *Value) ToString() (string, error) {
	raw, err := v.V()
	if err != nil {
		return "", err
	}

	s, ok := raw.(string)
	if !ok {
		return "", NotValidType
	}

	return s, nil
}

func (v *Value) MustString() string {
	s, _ := v.ToString()

	return s
}

func (v *Value) Compare(s *Scope, op Token, y *Value) bool {
	v.vari.scope = s
	y.vari.scope = s

	raw, err := v.V()
	if err != nil {
		fmt.Println(err, "<-====")
	}

	fmt.Println(v)

	switch raw.(type) {
	case int:
		i := v.MustInt()
		y, err := y.ToInt()
		if err != nil {
			fmt.Println(err)
			return false
		}

		fmt.Println(y, op, i)

		switch op {
		case Equals:
			return i == y
		case NotEquals:
			return i != y
		}

		fmt.Println("shouldnt be here")
		return false
	}

	return false
}

type variable struct {
	name  string
	scope *Scope
}
