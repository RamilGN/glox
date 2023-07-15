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
		{
			title:  "digit",
			source: []rune("523.23"),
			expected: []Token{
				{tType: number, lexeme: []rune("523.23"), literal: 523.23, line: 1},
				{tType: eof, lexeme: []rune{}, literal: nil, line: 1},
			},
		},
		{
			title:  "identifier",
			source: []rune("forbs"),
			expected: []Token{
				{tType: identifier, lexeme: []rune("forbs"), literal: nil, line: 1},
				{tType: eof, lexeme: []rune{}, literal: nil, line: 1},
			},
		},
		{
			title:  "reserved keyword",
			source: []rune("if"),
			expected: []Token{
				{tType: ifw, lexeme: []rune("if"), literal: nil, line: 1},
				{tType: eof, lexeme: []rune{}, literal: nil, line: 1},
			},
		},
		{
			title:  "multiline comment blocks",
			source: []rune("/*foo /* */ \n\nbar*/"),
			expected: []Token{
				{tType: eof, lexeme: []rune{}, literal: nil, line: 3},
			},
		},
	}

	for _, ts := range tests {
		t.Run(ts.title, func(t *testing.T) {
			s := NewScanner(ts.source)
			s.ScanTokens()
			assertEqualSlice(t, ts.expected, s.tokens)
		})
	}
}
