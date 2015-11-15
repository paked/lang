package lang

import (
	"bytes"
	"fmt"
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

func TestParsingNumber(t *testing.T) {
	src := `x int = 123`

	l := NewLexer(strings.NewReader(src))
	p := NewParser(l)

	prog := p.Parse()
	prog.Run()

	v := prog.scope.Get("x")
	if v == nil {
		t.Error("could not get x from scope")
		return
	}

	i, err := v.ToInt()
	if err != nil {
		t.Error("Coudl not move v to int")
	}

	if i != 123 {
		t.Errorf("wrong value, expected %v got %v", 123, i)
	}
}

func TestParsingIf(t *testing.T) {
	src := `if 9 == 9 {
x int = 22
y string = "ping"
print("ping")
}
`

	l := NewLexer(strings.NewReader(src))
	p := NewParser(l)

	prog := p.Parse()
	fmt.Println("===BEGIN===")
	prog.Run()
	fmt.Println("===END===")
}

func TestParsingBlock(t *testing.T) {
	src := `z int = 42
{
	x int = 22
	y string = "party!"
}`

	l := NewLexer(strings.NewReader(src))
	p := NewParser(l)

	prog := p.Parse()
	prog.Run()

	fmt.Println(prog.scope)

	v := prog.scope.Get("x")
	if v == nil {
		t.Error("could not get x from scope")
	}

	i, err := v.ToInt()
	if err != nil {
		t.Error("could nor cast to x")
	}

	if i != 22 {
		t.Error("wrong value in x")
	}

	v = prog.scope.Get("y")
	if v == nil {
		t.Error("could not get y froms cope")
	}

	str, err := v.ToString()
	if err != nil {
		t.Error("could not cast to string")
	}

	if str != "party!" {
		t.Error("wrong value in y")
	}

	v = prog.scope.Get("z")
	if v == nil {
		t.Error("could not get z")
	}

	i, err = v.ToInt()
	if err != nil {
		t.Error("could not cast z to int")
	}

	if i != 42 {
		t.Error("wrong value z")
	}
}
