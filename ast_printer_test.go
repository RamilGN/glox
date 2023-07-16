package main

import "testing"

func TestAstPrinterPrint(t *testing.T) {
	expression := Binary{
		Left: NewUnary(
			NewToken(minus, []rune("-"), nil, 1),
			NewLiteral(123),
		),
		Operator: NewToken(star, []rune("*"), nil, 1),
		Right:    NewGrouping(NewLiteral(45.67)),
	}

	astPrinter := new(AstPrinter)
	actual, err := astPrinter.Print(&expression)
	assertNil(t, err)
	assertEqualString(t, "(* (- 123) (group 45.67))", actual)
}
