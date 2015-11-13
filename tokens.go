package lang

type Token int

const (
	Illegal Token = iota + 1
	EOF
	Whitespace

	// Names
	Identifier

	// Chars
	Quotes
	OpenParen
	CloseParen

	// Assignment
	Assign

	// Types
	String
)
