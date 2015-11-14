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

func TestLexTypes(t *testing.T) {
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
