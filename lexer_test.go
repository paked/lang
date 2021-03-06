package lang

import (
	"fmt"
	"strings"
	"testing"
)

func TestLexWhitespace(t *testing.T) {
	src := `  `
	l := NewLexer(strings.NewReader(src))

	tok, s := l.Scan()
	if tok != Whitespace {
		t.Errorf("Whitespace not correct token, got %v expected %v", tok, Whitespace)
	}

	if s != src {
		t.Errorf("Wrong strings `%v` vs `%v`", s, src)
	}
}

func TestLexIdentifiers(t *testing.T) {
	src := `party`

	l := NewLexer(strings.NewReader(src))

	// test x
	tok, s := l.Scan()
	if tok != Identifier {
		t.Errorf("Wrong token name got %v, wanted %v", tok, Identifier)
	}

	if s != src {
		t.Errorf("Wrong literal got %v, wanted %v", s, src)
	}
}

func TestLexTypeInt(t *testing.T) {
	src := `message int`

	l := NewLexer(strings.NewReader(src))

	// test x
	tok, s := l.Scan()
	if tok != Identifier {
		t.Errorf("Wrong token name got %v, wanted %v", tok, Identifier)
	}

	if s != "message" {
		t.Errorf("Wrong literal got %v, wanted %v", s, "message")
	}

	// skip string
	l.Scan()

	// test string
	tok, s = l.Scan()
	if tok != Int {
		t.Errorf("(2) Expected String type got %v, wanted %v", tok, Int)
	}

	if s != "int" {
		t.Errorf("(2) Wrong naming literal got %v, wanted %v", s, "x")
	}
}

func TestLexValueInt(t *testing.T) {
	src := `message int = 123`

	l := NewLexer(strings.NewReader(src))

	// skip ident
	l.Scan()
	// skip whitespace
	l.Scan()
	// skip int
	l.Scan()
	// skip whitespace
	l.Scan()
	// skip '='
	l.Scan()
	// skip whitespace
	l.Scan()

	tok, lit := l.Scan()
	if tok != Number {
		t.Errorf("Wrong value type, got %v expected %v", tok, Number)
	}

	if lit != "123" {
		t.Errorf("Wrong value lit, go %v expected %v", lit, "123")
	}
}

func TestLexTypeString(t *testing.T) {
	src := `message string`

	l := NewLexer(strings.NewReader(src))

	// test x
	tok, s := l.Scan()
	if tok != Identifier {
		t.Errorf("Wrong token name got %v, wanted %v", tok, Identifier)
	}

	if s != "message" {
		t.Errorf("Wrong literal got %v, wanted %v", s, "message")
	}

	// skip string
	l.Scan()

	// test string
	tok, s = l.Scan()
	if tok != String {
		t.Errorf("(2) Expected String type got %v, wanted %v", tok, Identifier)
	}

	if s != "string" {
		t.Errorf("(2) Wrong naming literal got %v, wanted %v", s, "x")
	}
}

func TestAssignment(t *testing.T) {
	src := `msg string = _`

	l := NewLexer(strings.NewReader(src))

	// skip msg
	l.Scan()
	// skip whitespace
	l.Scan()
	// skip type
	l.Scan()
	// skip whitespace
	l.Scan()

	tok, s := l.Scan()
	if tok != Assign {
		t.Errorf("Wrong type for assignment got %v wanted %v", tok, Assign)
	}

	if s != "=" {
		t.Errorf("Wrong literal for assignment got %v expected %v", tok, Assign)
	}
}

func TestMultipleLines(t *testing.T) {
	src := `x string = "hey"
y string = "hello"`

	l := NewLexer(strings.NewReader(src))

	for {
		tok, lit := l.Scan()

		fmt.Println(tok, lit)

		if tok == EOF {
			break
		}
	}
}

func TestLexIf(t *testing.T) {
	src := `if 0 == 0 { print("hello") }`

	l := NewLexer(strings.NewReader(src))

	tok, _ := l.Scan()
	if tok != If {
		t.Errorf("expected if got %v", tok)
	}

	l.Scan()
	l.Scan()
	l.Scan()

	tok, _ = l.Scan()
	if tok != Equals {
		t.Errorf("expected equals got %v", tok)
	}

	l.Scan()
	l.Scan()
	l.Scan()

	tok, _ = l.Scan()
	if tok != OpenBracket {
		t.Errorf("expected open bracket got %v", tok)
	}
}
