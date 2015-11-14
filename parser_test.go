package lang

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	x := "heyy how are you"
	src := `x string = "` + x + `"`

	l := NewLexer(strings.NewReader(src))

	p := NewParser(l)

	prog := p.Parse()

	prog.Run()

	if xv := prog.scope.Get("x"); xv != x {
		t.Errorf("wrong value for x... got '%v' expected '%v'", xv, x)
	}
}
