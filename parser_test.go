package lang

import (
	"bytes"
	"strings"
	"testing"
)

func TestParsingAssignment(t *testing.T) {
	x := "heyy how are you"
	src := `x string = "` + x + `"`

	l := NewLexer(strings.NewReader(src))

	p := NewParser(l)

	prog := p.Parse()

	prog.Run()

	if xv := prog.scope.Get("x").MustString(); xv != x {
		t.Errorf("wrong value for x... got '%v' expected '%v'", xv, x)
	}
}

func TestParsingMultilineAssignment(t *testing.T) {
	src := `x string = "hello"
y string = "no no"`

	l := NewLexer(strings.NewReader(src))
	p := NewParser(l)

	prog := p.Parse()
	prog.Run()

	if v := prog.scope.Get("x").MustString(); v != "hello" {
		t.Error("wrong value for x... got '%v' expected 'hello'", v)
	}

	if v := prog.scope.Get("y").MustString(); v != "no no" {
		t.Error("wrong value for y... got '%v' expected 'no no'", v)
	}
}

func TestParsingFunction(t *testing.T) {
	src := `print("hello")`

	l := NewLexer(strings.NewReader(src))
	p := NewParser(l)

	var out bytes.Buffer

	prog := p.Parse()

	prog.out = &out
	prog.Run()

	if out.String() != "hello" {
		t.Errorf("wrong value for print got '%v' wanted 'hello'", out.String())
	}
}
