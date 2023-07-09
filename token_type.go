package main

//go:generate stringer --type TokenType --output=./token_type_stringer.gen.go

type TokenType int

const (
	// Single-character tokens.
	leftParen TokenType = iota
	rightParen
	leftBrace
	rightBrace
	comma
	dot
	minus
	plus
	semicolon
	slash
	star

	// One or two character tokens.
	bang
	bangEqual
	equal
	equalEqual
	greater
	greaterEqual
	less
	lessEqual

	// Literals.
	identifier
	stringw
	number

	// Keywords.
	and
	class
	elsew
	falsew
	fun
	forw
	ifw
	nilw
	or
	printw
	returnw
	super
	this
	truew
	varw
	while

	eof
)

func getReservedKeywords() map[string]TokenType {
	return map[string]TokenType{
		"and":    and,
		"class":  class,
		"else":   elsew,
		"false":  falsew,
		"for":    forw,
		"fun":    fun,
		"if":     ifw,
		"nil":    nilw,
		"or":     or,
		"print":  printw,
		"return": returnw,
		"super":  super,
		"this":   this,
		"true":   truew,
		"var":    varw,
		"while":  while,
	}
}
