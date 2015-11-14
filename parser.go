package lang

import (
	"errors"
	"fmt"
	"os"
)

type Parser struct {
	l *Lexer

	buf []buf
	n   int
}

func NewParser(l *Lexer) *Parser {
	return &Parser{
		l: l,
	}
}

func (p *Parser) Parse() *Program {
	prog := &Program{
		scope: &Scope{
			values: make(map[string]*Value),
		},
		out: os.Stdout,
	}

	for {
		n := p.n
		tok, lit := p.scanSkipWhitespace()
		if tok == EOF {
			fmt.Println("REACHED EOF!!!")
			break
		}

		p.unscan()

		if p.is(MatchAssignment...) {
			as, err := p.parseAssignment()
			if err == nil {
				prog.statements = append(prog.statements, as)
			} else {
				p.reset(n)
			}
		} else if p.is(Identifier, OpenParen) {
			f, err := p.parseFunction()
			if err == nil {
				prog.statements = append(prog.statements, f)
			} else {
				p.reset(n)
			}
		}

		tok, lit = p.scan()
		if tok == Whitespace && lit == "\n" {
			fmt.Println("reached termination!")
			continue
		}

		p.unscan()

		fmt.Println("didnt match: SYNTAX ERROR")
	}

	return prog
}

func (p *Parser) parseFunction() (*FunctionStatement, error) {
	f := &FunctionStatement{}
	tok, lit := p.scan()

	if tok != Identifier {
		return nil, fmt.Errorf("not correct syntax")
	}

	f.Name = lit

	// skip opening paren
	p.scan()

	var params string

	tok, lit = p.scan()
	if tok != Quotes {
		return nil, fmt.Errorf("well that didnt work")
	}

	for {
		tok, lit = p.scan()
		if tok == EOF {
			return nil, errors.New("eof")
		}

		if tok == Quotes {
			p.unscan()
			break
		}

		params += lit
	}

	tok, lit = p.scan()
	if tok != Quotes {
		return nil, fmt.Errorf("well that didnt work")
	}

	f.Params = params

	// skip closing paren
	p.scan()

	return f, nil
}

var MatchAssignment = []Token{Identifier, Whitespace, String}

func (p *Parser) parseAssignment() (*AssignmentStatement, error) {
	tok, lit := p.scanSkipWhitespace()
	// Got name
	assign := &AssignmentStatement{
		Name: lit,
	}

	tok, lit = p.scanSkipWhitespace()
	if tok != String {
		fmt.Println("NOT TYPE")
		return nil, fmt.Errorf("found %v expected String")
	}

	tok, lit = p.scanSkipWhitespace()
	if tok != Assign {
		fmt.Println("NOT ASSIGN. TIME TO DIE!")
		return nil, fmt.Errorf("found %v expected String")
	}

	tok, lit = p.scanSkipWhitespace()
	if tok != Quotes {
		fmt.Println("NOT quotes. TIME TO DIE")
		fmt.Println(tok, lit)

		return nil, fmt.Errorf("found %v expected quotes")
	}

	var buf string

	for {
		tok, lit := p.scan()

		if tok == EOF {
			p.unscan()
			break
		}

		if tok == Quotes {
			break
		}

		buf += lit
	}

	fmt.Println("[DONE] got value", buf)

	v, err := NewValue(buf)
	if err != nil {
		return nil, err
	}

	assign.Value = v

	return assign, nil
}

func (p *Parser) is(ts ...Token) bool {
	for _, t := range ts {
		tok, _ := p.scan()

		defer func() { p.unscan() }()

		if tok != t {
			fmt.Println("Got", t, "expected", tok)
			return false
		}
	}

	return true
}

// If it can pull n from tokens, do that... else scan new tok
// and add it to the buf
func (p *Parser) scan() (Token, string) {
	defer func() {
		p.n++
	}()

	if p.n >= len(p.buf) {
		tok, lit := p.l.Scan()
		fmt.Println("Scanning new token", p.n, tok)
		p.buf = append(p.buf, buf{
			tok: tok,
			lit: lit,
		})

		return tok, lit
	}

	b := p.buf[p.n]

	fmt.Println("Retrieving old token", p.n, b.tok)
	return b.tok, b.lit
}

func (p *Parser) unscan() {
	p.n--
}

func (p *Parser) scanSkipWhitespace() (Token, string) {
	for {
		tok, lit := p.scan()

		if tok != Whitespace {
			return tok, lit
		}
	}
}

func (p *Parser) reset(n int) {
	for p.n > n {
		p.unscan()
	}
}

type buf struct {
	tok Token
	lit string
}