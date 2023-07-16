package main

import "fmt"

type Token struct {
	tType   TokenType
	lexeme  []rune
	literal any
	line    int
}

func NewToken(tType TokenType, lexeme []rune, literal any, line int) Token {
	return Token{
		tType:   tType,
		lexeme:  lexeme,
		literal: literal,
		line:    line,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %v", t.tType, string(t.lexeme), t.literal)
}
