package main

import "fmt"

type Token struct {
	tType   TokenType
	lexeme  []rune
	literal any
	line    int
}

func (t *Token) String() string {
	return fmt.Sprintf("%d %v %s", t.tType, t.lexeme, t.literal)
}
