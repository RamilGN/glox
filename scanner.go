package main

import (
	"strconv"
)

type Scanner struct {
	source []rune
	tokens []Token

	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens() {
	s.line++
	s.start = 0
	s.current = 0

	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{tType: eof, lexeme: []rune{}, literal: nil, line: s.line})
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(leftParen, nil)
	case ')':
		s.addToken(rightParen, nil)
	case '{':
		s.addToken(leftBrace, nil)
	case '}':
		s.addToken(rightBrace, nil)
	case ',':
		s.addToken(comma, nil)
	case '.':
		s.addToken(dot, nil)
	case '-':
		s.addToken(minus, nil)
	case '+':
		s.addToken(plus, nil)
	case ';':
		s.addToken(semicolon, nil)
	case '*':
		s.addToken(star, nil)
	case '!':
		var t TokenType
		if s.match('=') {
			t = bangEqual
		} else {
			t = bang
		}

		s.addToken(t, nil)
	case '=':
		var t TokenType
		if s.match('=') {
			t = equalEqual
		} else {
			t = equal
		}

		s.addToken(t, nil)
	case '<':
		var t TokenType
		if s.match('=') {
			t = lessEqual
		} else {
			t = less
		}

		s.addToken(t, nil)
	case '>':
		var t TokenType
		if s.match('=') {
			t = greaterEqual
		} else {
			t = greater
		}

		s.addToken(t, nil)
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(slash, nil)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	case '"':
		s.scanString()
	default:
		switch {
		case s.isDigit(c):
			s.scanNumber()
		case s.isAlpha(c):
			s.scanIdentifier()
		default:
			errorm(s.line, "Unexpected character.")
		}
	}
}

func (s *Scanner) advance() rune {
	char := s.source[s.current]
	s.current++

	return char
}

func (s *Scanner) addToken(tType TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{tType: tType, lexeme: text, literal: literal, line: s.line})
}

// match consumes next expected char.
func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.advance()

	return true
}

// cur gets current rune from source.
func (s *Scanner) cur() rune {
	return s.source[s.current]
}

// Peek is safe cur.
func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}

	return s.cur()
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\000'
	}

	return s.source[s.current+1]
}

// isAtEnd checks if the scanner can't scan next token.
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// isDigit checks if rune is digit.
func (s *Scanner) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// isAlpha checks if rune is char.
func (s *Scanner) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

// isAlpha checks if rune is char or digit.
func (s *Scanner) isAlphaNumeric(c rune) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

// scanString scan string with double quotes.
func (s *Scanner) scanString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.cur() == '\n' {
			s.line++
		}

		s.advance()
	}

	if s.isAtEnd() {
		errorm(s.line, "Unterminated string")
		return
	}

	s.advance()

	text := s.source[s.start+1 : s.current-1] // trim the surronding quotes
	s.addToken(stringw, text)
}

func (s *Scanner) scanNumber() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	num, _ := strconv.ParseFloat(string(s.source[s.start:s.current]), 64)
	s.addToken(number, num)
}

func (s *Scanner) scanIdentifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]

	tType, ok := reservedKeywords[string(text)]
	if !ok {
		tType = identifier
	}

	s.addToken(tType, nil)
}
