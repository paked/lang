package lang

import "fmt"

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
			make(map[string]string),
		},
	}

	for {
		tok, lit := p.scanSkipWhitespace()
		if tok == EOF {
			fmt.Println("REACHED EOF!!!")
			break
		}

		switch tok {
		case Identifier:
			// Got name
			assign := &AssignmentStatement{
				Name: lit,
			}

			tok, lit = p.scanSkipWhitespace()
			if tok != String {
				fmt.Println("NOT TYPE")
				fmt.Println(tok, lit)
				break
			}

			tok, lit = p.scanSkipWhitespace()
			if tok != Assign {
				fmt.Println("NOT ASSIGN. TIME TO DIE!")
				break
			}

			tok, lit = p.scanSkipWhitespace()
			if tok != Quotes {
				fmt.Println("NOT quotes. TIME TO DIE")
				fmt.Println(tok, lit)
				break
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

			assign.Value = buf

			prog.statements = append(prog.statements, assign)

			continue
		}
	}

	fmt.Println(prog)

	return prog
}

func (p *Parser) scan() (Token, string) {
	p.n++

	// check if n is in p.buf (n < len p.buf)
	if p.n < len(p.buf) {
		fmt.Println("pulling from buffer")
		b := p.buf[len(p.buf)-1]
		return b.tok, b.lit
	}

	tok, lit := p.l.Scan()
	p.buf = append(p.buf, buf{
		tok: tok,
		lit: lit,
	})

	return tok, lit
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

type buf struct {
	tok Token
	lit string
}
