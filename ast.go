package lang

import (
	"errors"
	"fmt"
	"io"
)

type Scope struct {
	values map[string]*Value
	out    io.Writer
}

func (s *Scope) Set(key string, val *Value) {
	s.values[key] = val
}

func (s *Scope) Get(key string) *Value {
	return s.values[key]
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
	Value *Value
}

func (as *AssignmentStatement) Eval(s *Scope) error {
	s.Set(as.Name, as.Value)

	return nil
}

func (as *AssignmentStatement) String() string {
	return fmt.Sprintf("%v = %v", as.Name, as.Value.v)
}

type FunctionStatement struct {
	Name   string
	Params string
}

func (f *FunctionStatement) Eval(s *Scope) error {
	// TODO: pull function from scope
	if f.Name == "print" {
		fmt.Fprint(s.out, f.Params)
	}

	return nil
}

type IfStatement struct {
	A *Value
	B *Value

	Op Token

	Then *BlockStatement
}

func (is *IfStatement) Eval(s *Scope) error {
	if is.A.Compare(is.Op, is.B) {
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
	for _, stmt := range is.Statements {
		err := stmt.Eval(s)
		if err != nil {
			return err
		}
	}

	return nil
}
