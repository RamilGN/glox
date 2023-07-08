package main

import "fmt"

type Token struct {
	literal any
	lexeme  []rune
	line    int
	tType   TokenType
}

func (t *Token) String() string {
	return fmt.Sprintf("%d %v %s", t.tType, t.lexeme, t.literal)
}
