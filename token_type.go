package main

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
