package main

type Scanner struct {
	source []rune
	tokens []Token

	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens() {
	s.line += 1
	s.start = 0
	s.current = 0

	for s.current < len(s.source) {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{tType: eof, lexeme: []rune{}, literal: nil, line: s.line})
}

func (s *Scanner) scanToken() {
	switch s.cur() {
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
				s.current++
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
		s.current++

		for s.peek() != '"' && !s.isAtEnd() {
			if s.cur() == '\n' {
				s.line++
			}

			s.current++
		}

		if s.source[s.current] != '"' {
			errorm(s.line, "Unterminated string")
			break
		}

		val := s.source[s.start+1 : s.current] // trim the surronding quotes
		s.addToken(stringw, val)
	default:
		errorm(s.line, "Unexpected character.")
	}

	s.current++
}

func (s *Scanner) addToken(tType TokenType, literal any) {
	text := s.source[s.start : s.current+1]
	s.tokens = append(s.tokens, Token{tType: tType, lexeme: text, literal: literal, line: s.line})
}

// match consumes next expected char
func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current+1] != expected {
		return false
	}

	s.current++
	return true
}

// cur gets current rune from source
func (s *Scanner) cur() rune {
	return s.source[s.current]
}

// peek is safe cur
func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}

	return s.cur()
}

// isAtEnd checks if the scanner on the last element
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)-1
}
