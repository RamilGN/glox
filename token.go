package main

import "fmt"

type Token struct {
	literal any
	lexeme  []rune
	line    int
	tType   TokenType
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", t.tType, string(t.lexeme), t.literal)
}
