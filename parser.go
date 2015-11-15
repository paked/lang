package lang

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
			break
		}

		p.unscan()

		stmt, err := p.parseStatement()
		if err == nil {
			prog.statements = append(prog.statements, stmt)
		} else {
			p.reset(n)
		}

		tok, lit = p.scan()
		if tok == Whitespace && lit == "\n" {
			continue
		} else if tok == EOF {
			break
		}

		p.unscan()

		fmt.Println("didnt match: SYNTAX ERROR")
		break
	}

	return prog
}

func (p *Parser) parseStatement() (Statement, error) {
	if p.is(MatchStringAssignment...) || p.is(MatchIntAssignment...) {
		return p.parseAssignment()
	} else if p.is(Identifier, OpenParen) {
		return p.parseFunction()
	} else if p.is(If, Whitespace, Any, Whitespace, Any, Whitespace) {
		return p.parseIf()
	} else if p.is(OpenBracket) {
		return p.parseBlock()
	}

	return nil, errors.New("no statement found")
}

func (p *Parser) parseBlock() (*BlockStatement, error) {
	n := p.n
	bs := &BlockStatement{}

	tok, _ := p.scan()
	if tok != OpenBracket {
		return nil, errors.New("expected open bracket")
	}

	for {
		tok, _ := p.scanSkipWhitespace()

		if tok == EOF {
			p.unscan()

			return nil, errors.New("unexpected EOF")
		}

		if tok == CloseBracket {
			break
		}

		p.unscan()

		stmt, err := p.parseStatement()
		if err == nil {
			bs.Statements = append(bs.Statements, stmt)
		} else {
			p.reset(n)
			fmt.Println("COULD NOT PARSE STATEMENT")
			return nil, errors.New("THAT FAILED")
		}
	}

	return bs, nil
}

func (p *Parser) parseIf() (*IfStatement, error) {
	tok, _ := p.scan()
	if tok != If {
		return nil, errors.New("this is an error")
	}

	tok, _ = p.scanSkipWhitespace()
	p.unscan()

	a, err := p.parseLiteral()
	if err != nil {
		return nil, err
	}

	tok, _ = p.scanSkipWhitespace()
	if tok != Equals && tok != NotEquals {
		return nil, errors.New("not implemtned")
	}

	op := tok

	p.scanSkipWhitespace()
	p.unscan()

	b, err := p.parseLiteral()
	if err != nil {
		return nil, err
	}

	p.scan()

	block, err := p.parseBlock()
	if err != nil {
		fmt.Println("invalid block", err)
		return nil, err
	}

	fmt.Println("<====>")
	fmt.Println(block)

	is := &IfStatement{
		A:    a,
		B:    b,
		Op:   op,
		Then: block,
	}

	return is, nil
}

func (p *Parser) parseString() (string, error) {
	var s string

	tok, lit := p.scan()
	if tok != Quotes {
		return "", fmt.Errorf("expected quotes got: %v (%v)", tok, lit)
	}

	for {
		tok, lit = p.scan()
		if tok == EOF {
			p.unscan()
			return "", errors.New("eof")
		}

		if tok == Quotes {
			break
		}

		s += lit
	}

	return s, nil
}

func (p *Parser) parseNumber() (int, error) {
	tok, lit := p.scan()
	if tok != Number {
		return 0, fmt.Errorf("Wrong token type expected %v got %v(%v)", Number, tok, lit)
	}

	i, err := strconv.Atoi(lit)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (p *Parser) parseLiteral() (*Value, error) {
	n := p.n

	s, err := p.parseString()
	if err == nil {
		return NewValue(s)
	}

	p.reset(n)

	i, err := p.parseNumber()
	if err == nil {
		return NewValue(i)
	}

	p.reset(n)

	return nil, errors.New("no literal")
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

	v, err := p.parseLiteral()
	if err != nil {
		return nil, err
	}

	f.Params = v.MustString()

	// skip closing paren
	p.scan()

	return f, nil
}

var MatchStringAssignment = []Token{Identifier, Whitespace, String}
var MatchIntAssignment = []Token{Identifier, Whitespace, Int}

func (p *Parser) parseAssignment() (*AssignmentStatement, error) {
	tok, lit := p.scanSkipWhitespace()
	// Got name
	assign := &AssignmentStatement{
		Name: lit,
	}

	tok, lit = p.scanSkipWhitespace()
	if tok != String && tok != Int {
		return nil, fmt.Errorf("found %v expected String or Int")
	}

	tok, lit = p.scanSkipWhitespace()
	if tok != Assign {
		return nil, fmt.Errorf("found %v expected Assign")
	}

	tok, lit = p.scan()

	v, err := p.parseLiteral()
	if err != nil {
		return nil, err
	}

	assign.Value = v

	fmt.Println(assign)

	return assign, nil
}

func (p *Parser) is(ts ...Token) bool {
	for _, t := range ts {
		tok, _ := p.scan()

		defer func() { p.unscan() }()

		if t == Any {
			continue
		}

		if tok != t {
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
		p.buf = append(p.buf, buf{
			tok: tok,
			lit: lit,
		})

		return tok, lit
	}

	b := p.buf[p.n]

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
