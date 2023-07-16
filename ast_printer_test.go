package main

import "testing"

func TestAstPrinterPrint(t *testing.T) {
	expression := Binary{
		Left: &Unary{
			Operator: Token{tType: minus, lexeme: []rune("-"), literal: nil, line: 1},
			Right:    &Literal{Object: 123},
		},
		Operator: Token{tType: star, lexeme: []rune("*"), literal: nil, line: 1},
		Right:    &Grouping{Expression: &Literal{Object: 45.67}},
	}

	astPrinter := new(AstPrinter)
	actual, err := astPrinter.Print(&expression)
	assertNil(t, err)
	assertEqualString(t, "(* (- 123) (group 45.67))", actual)
}
