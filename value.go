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

type Value interface {
	Lit(*Scope) (*Literal, error)
}

type Literal struct {
	typ ValueType
	raw interface{}
}

func NewLiteral(raw interface{}) (*Literal, error) {
	lit := &Literal{
		raw: raw,
	}

	switch raw.(type) {
	case int:
		lit.typ = ValueInt
	case string:
		lit.typ = ValueString
	default:
		return nil, errors.New("couldnt creater lit")
	}

	return lit, nil
}

func (lit *Literal) String() string {
	return fmt.Sprint(lit.raw)
}

func (lit *Literal) Lit(*Scope) (*Literal, error) {
	return lit, nil
}

func (lit *Literal) ToInt() (int, error) {
	if lit.typ != ValueInt {
		return 0, errors.New("failed cast")
	}

	return lit.raw.(int), nil
}

func (lit *Literal) MustInt() int {
	i, _ := lit.ToInt()

	return i
}

func (lit *Literal) ToString() (string, error) {
	if lit.typ != ValueString {
		return "", errors.New("failed value")
	}

	return lit.raw.(string), nil
}

func (lit *Literal) Compare(op Token, b *Literal) bool {
	if lit.typ != b.typ {
		return false
	}

	switch lit.typ {
	case ValueInt:
		x := lit.MustInt()
		y := lit.MustInt()

		switch op {
		case Equals:
			return x == y
		case NotEquals:
			return x != y
		}
	case ValueString:
		x := lit.MustString()
		y := lit.MustString()

		switch op {
		case Equals:
			return x == y
		case NotEquals:
			return x != y
		}
	}

	return false
}

func (lit *Literal) MustString() string {
	s, _ := lit.ToString()

	return s
}

type Variable struct {
	name string
}

func NewVariable(name string) (*Variable, error) {
	return &Variable{
		name: name,
	}, nil
}

func (v *Variable) Lit(s *Scope) (*Literal, error) {
	lit := s.Get(v.name)
	if lit == nil {
		return lit, errors.New("couldnt get lit")
	}

	return lit, nil
}
