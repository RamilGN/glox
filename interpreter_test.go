package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestInterpreter(t *testing.T) {
	rw := bufio.NewReadWriter(bufio.NewReader(&bytes.Buffer{}), bufio.NewWriter(&bytes.Buffer{}))
	interpreter := Interpreter{lox: NewLox(rw)}

	tests := []struct {
		title      string
		expression Expr
		expected   any
	}{
		{
			title:      "grouping",
			expression: NewGrouping(NewLiteral(43.23)),
			expected:   43.23,
		},
		{
			title:      "literal",
			expression: NewLiteral("bar"),
			expected:   "bar",
		},
		// Unaries.
		{
			title: "unary minus",
			expression: NewUnary(
				NewToken(minus, []rune("-"), nil, 1),
				NewLiteral(34),
			),
			expected: -34.0,
		},
		{
			title: "unary bang with false",
			expression: NewUnary(
				NewToken(bang, []rune("!"), nil, 1),
				NewLiteral(false),
			),
			expected: true,
		},
		{
			title: "unary bang with nil",
			expression: NewUnary(
				NewToken(bang, []rune("!"), nil, 1),
				NewLiteral(nil),
			),
			expected: true,
		},
		{
			title: "unary bang with true",
			expression: NewUnary(
				NewToken(bang, []rune("!"), nil, 1),
				NewLiteral(true),
			),
			expected: false,
		},
		// Binaries
		{
			title: "binary minus",
			expression: NewBinary(
				NewLiteral(45.23),
				NewToken(minus, []rune("-"), nil, 1),
				NewLiteral(45.23),
			),
			expected: 0.0,
		},
		{
			title: "binary plus on nums",
			expression: NewBinary(
				NewLiteral(34),
				NewToken(plus, []rune("+"), nil, 1),
				NewLiteral(45.23),
			),
			expected: 79.23,
		},
		{
			title: "binary plus on strings",
			expression: NewBinary(
				NewLiteral("Hello, "),
				NewToken(plus, []rune("+"), nil, 1),
				NewLiteral("World!"),
			),
			expected: "Hello, World!",
		},
		// Errors
	}

	for _, v := range tests {
		t.Run(v.title, func(t *testing.T) {
			actual, err := interpreter.Interpret(v.expression)
			requireNil(t, err)
			assertEqual(t, v.expected, actual)
		})
	}

	errTests := []struct {
		title      string
		expression Expr
		expected   string
	}{
		{
			title: "err. unary on string",
			expression: NewUnary(
				NewToken(minus, []rune("-"), nil, 1),
				NewLiteral("foo"),
			),
			expected: "RuntimeError: operand must be a number `-foo` | minus - <nil>",
		},
		{
			title: "err. greater on string",
			expression: NewBinary(
				NewLiteral("foo"),
				NewToken(greater, []rune(">"), nil, 1),
				NewLiteral(3),
			),
			expected: "RuntimeError: operands must be numbers. `foo > 3` | greater > <nil>",
		},
		{
			title: "err. adding string to num",
			expression: NewBinary(
				NewLiteral("foo"),
				NewToken(plus, []rune("+"), nil, 1),
				NewLiteral(3),
			),
			expected: "RuntimeError: addition error. `foo + 3` | plus + <nil>",
		},
	}

	for _, v := range errTests {
		t.Run(v.title, func(t *testing.T) {
			_, err := interpreter.Interpret(v.expression)
			requireNotNil(t, err)
			assertEqualString(t, v.expected, err.Error())
		})
	}
}
