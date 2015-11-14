package lang

//go:generate stringer -type=Token
type Token int

const (
	Illegal Token = iota + 1
	Any
	EOF
	Whitespace

	// Names
	Identifier

	// Chars
	Quotes
	OpenParen
	CloseParen
	OpenBracket
	CloseBracket

	// Assignment
	Assign

	// Types
	String
	Int

	// Vals
	Number

	// Equality
	Equals
	NotEquals

	// Keywords
	If
)
