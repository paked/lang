package lang

import "errors"

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
)

type Value struct {
	typ ValueType
	v   interface{}
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
	}

	return nil, errors.New("Not implemented")
}

func (v *Value) ToInt() (int, error) {
	i, ok := v.v.(int)
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
	s, ok := v.v.(string)
	if !ok {
		return "", NotValidType
	}

	return s, nil
}

func (v *Value) MustString() string {
	s, _ := v.ToString()

	return s
}
