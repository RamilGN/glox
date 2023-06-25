package main

import (
	"testing"
)

func TestScanner(t *testing.T) {
	tests := []struct {
		title    string
		source   []rune
		expected []Token
	}{
		{
			title:  "Single tokens",
			source: []rune("(}*"),
			expected: []Token{
				{tType: leftParen, lexeme: []rune{'('}, literal: nil, line: 1},
				{tType: rightBrace, lexeme: []rune{'}'}, literal: nil, line: 1},
				{tType: star, lexeme: []rune{'*'}, literal: nil, line: 1},
				{tType: eof, lexeme: []rune{}, literal: nil, line: 1},
			},
		},
		{
			title:  "Double tokens with single token",
			source: []rune("!=!"),
			expected: []Token{
				{tType: bangEqual, lexeme: []rune("!="), literal: nil, line: 1},
				{tType: bang, lexeme: []rune("!"), literal: nil, line: 1},
				{tType: eof, lexeme: []rune{}, literal: nil, line: 1},
			},
		},
		{
			title:  "Empty",
			source: []rune(""),
			expected: []Token{
				{tType: eof, lexeme: []rune{}, literal: nil, line: 1},
			},
		},
		{
			title:  "Strange tokens @",
			source: []rune(""),
			expected: []Token{
				{tType: eof, lexeme: []rune{}, literal: nil, line: 1},
			},
		},
		{
			title:  "Comments",
			source: []rune("// comment line"),
			expected: []Token{
				{tType: eof, lexeme: []rune{}, literal: nil, line: 1},
			},
		},
		{
			title:  "New line",
			source: []rune("-\n\n+"),
			expected: []Token{
				{tType: minus, lexeme: []rune("-"), literal: nil, line: 1},
				{tType: plus, lexeme: []rune("+"), literal: nil, line: 3},
				{tType: eof, lexeme: []rune{}, literal: nil, line: 3},
			},
		},
		{
			title:  "multiline string",
			source: []rune("\"he\nwo\""),
			expected: []Token{
				{tType: stringw, lexeme: []rune("\"he\nwo\""), literal: []rune("he\nwo"), line: 2},
				{tType: eof, lexeme: []rune{}, literal: nil, line: 2},
			},
		},
		{
			title:  "unterminated string",
			source: []rune("\"he"),
			expected: []Token{
				{tType: eof, lexeme: []rune{}, literal: nil, line: 1},
			},
		},
	}

	for _, ts := range tests {
		t.Run(ts.title, func(t *testing.T) {
			s := Scanner{source: ts.source}
			s.ScanTokens()
			assertEqualSlice(t, ts.expected, s.tokens)
		})
	}
}
