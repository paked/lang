// Package lang is a simple programming language.
package lang

import (
	"bufio"
	"bytes"
	"io"
)

type Lexer struct {
	r *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		r: bufio.NewReader(r),
	}
}

func (l *Lexer) Scan() (Token, string) {
	ch := l.read()

	if l.isWhitespace(ch) {
		l.unread()

		return l.scanWhitespace()
	} else if l.isLetter(ch) {
		l.unread()

		return l.scanIdentifier()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '(':
		return OpenParen, "("
	case ')':
		return CloseParen, ")"
	case '=':
		return Assign, "="
	case '"':
		return Quotes, "\""
	}

	return Illegal, string(ch)
}

func (l *Lexer) isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (l *Lexer) scanWhitespace() (Token, string) {
	var buf bytes.Buffer

	for {
		ch := l.read()

		if ch == eof {
			break
		} else if !l.isWhitespace(ch) {
			l.unread()
			break
		}

		buf.WriteRune(ch)
	}

	return Whitespace, buf.String()
}

func (l *Lexer) isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (l *Lexer) isNumber(ch rune) bool {
	return ch > '0' && ch < '9'
}

func (l *Lexer) scanIdentifier() (Token, string) {
	var buf bytes.Buffer

	for {
		ch := l.read()

		if ch == eof {
			break
		} else if !l.isLetter(ch) {
			l.unread()
			break
		}

		buf.WriteRune(ch)
	}

	switch buf.String() {
	case "string":
		return String, buf.String()
	case "int":
		return Int, buf.String()
	}

	return Identifier, buf.String()
}

func (l *Lexer) read() rune {
	ch, _, err := l.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

func (l *Lexer) unread() {
	l.r.UnreadRune()
}

var eof = rune(0)
