package lang

import (
	"errors"
	"fmt"
	"io"
)

type Scope struct {
	parent *Scope
	values map[string]*Literal
	out    io.Writer
}

func (s *Scope) Set(key string, val *Literal) {
	s.values[key] = val
}

func (s *Scope) Get(key string) *Literal {
	lit := s.values[key]
	if lit == nil && s.parent != nil {
		return s.parent.Get(key)
	}

	return lit
}

func (s *Scope) Sub() *Scope {
	return &Scope{
		out:    s.out,
		values: make(map[string]*Literal),
		parent: s,
	}
}

type Program struct {
	scope      *Scope
	statements []Statement
	out        io.Writer
}

func (prog *Program) Run() error {
	prog.scope.out = prog.out
	for _, stmt := range prog.statements {
		stmt.Eval(prog.scope)
	}

	return nil
}

type Statement interface {
	Eval(*Scope) error
}

type AssignmentStatement struct {
	Name  string
	Type  string
	Value Value
}

func (as *AssignmentStatement) Eval(s *Scope) error {
	v, err := as.Value.Lit(s)
	if err != nil {
		return err
	}

	s.Set(as.Name, v)

	return nil
}

func (as *AssignmentStatement) String() string {
	return fmt.Sprintf("%v = %v", as.Name, as.Value)
}

type FunctionStatement struct {
	Name   string
	Params Value
}

func (f *FunctionStatement) Eval(s *Scope) error {
	// TODO: pull function from scope
	if f.Name == "print" {
		lit, err := f.Params.Lit(s)
		if err != nil {
			return err
		}

		fmt.Fprint(s.out, lit)
	}

	return nil
}

type IfStatement struct {
	A Value
	B Value

	Op Token

	Then *BlockStatement
}

func (is *IfStatement) Eval(s *Scope) error {
	a, err := is.A.Lit(s)
	if err != nil {
		return err
	}

	b, err := is.B.Lit(s)
	if err != nil {
		return err
	}

	if a.Compare(is.Op, b) {
		return is.Then.Eval(s)
	} else {
		fmt.Println("NOTICE: TEY ARE NOT EQUAL")
	}

	return errors.New("not implemented")
}

type BlockStatement struct {
	Statements []Statement
}

func (is *BlockStatement) Eval(s *Scope) error {
	scop := s.Sub()
	for _, stmt := range is.Statements {
		err := stmt.Eval(scop)
		if err != nil {
			return err
		}
	}

	return nil
}

type SetStatement struct {
	Name  string
	Value Value
}

func (ss *SetStatement) Eval(s *Scope) error {
	lit, err := ss.Value.Lit(s)
	if err != nil {
		return err
	}

	s.Set(ss.Name, lit)

	return nil
}
