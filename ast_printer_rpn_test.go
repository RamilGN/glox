package main

import "testing"

func TestAstPrinterRPNPrint(t *testing.T) {
	expression := Binary{
		Left: NewBinary(
			NewLiteral(1),
			NewToken(plus, []rune("+"), nil, 1),
			NewLiteral(2),
		),
		Operator: NewToken(star, []rune("*"), nil, 1),
		Right: NewBinary(
			NewLiteral(4),
			NewToken(minus, []rune("-"), nil, 1),
			NewLiteral(3),
		),
	}

	astPrinter := new(AstPrinterRPN)
	actual, err := astPrinter.Print(&expression)
	assertNil(t, err)
	assertEqualString(t, "1 2 + 4 3 - *", actual)
}
