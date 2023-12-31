//go:generate go run tools/ast/main.go

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var errLox = errors.New("[Lox]")

func errorm(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Fprintf(os.Stdin, "[line %d] Error %s: %s\n", line, where, message)
}

type Lox struct {
	rw              *bufio.ReadWriter
	hadError        bool
	hadRuntimeError bool
	interpreter     *Interpreter
}

func NewLox(rw *bufio.ReadWriter) *Lox {
	if rw == nil {
		rw = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	}

	l := Lox{rw: rw}
	i := Interpreter{lox: &l}
	l.interpreter = &i

	return &l
}

func (l *Lox) Start(args []string) error {
	switch len(args) {
	case 1:
		err := l.runFile(args[0])
		if err != nil {
			return fmt.Errorf("%w. %w", errLox, err)
		}

	case 0:
		err := l.runPrompt()
		if err != nil {
			return fmt.Errorf("%w. %w", errLox, err)
		}
	default:
		return fmt.Errorf("%w. usage: glox [filename]", errLox)
	}

	return nil
}

func (l *Lox) run(src string) error {
	scanner := NewScanner([]rune(src))
	scanner.ScanTokens()

	for _, v := range scanner.tokens {
		_, err := l.rw.WriteString(v.String())
		if err != nil {
			return fmt.Errorf("%w. %w", errLox, err)
		}

		_, err = l.rw.WriteString("\n")
		if err != nil {
			return fmt.Errorf("%w. %w", errLox, err)
		}

		err = l.rw.Flush()
		if err != nil {
			return fmt.Errorf("%w. %w", errLox, err)
		}
	}

	return nil
}

func (l *Lox) runFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("%w. %w", errLox, err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("%w. %w", errLox, err)
	}

	err = l.run(string(content))
	if err != nil {
		return fmt.Errorf("%w. %w", errLox, err)
	}

	if l.hadError {
		return fmt.Errorf("%w. %w", errLox, err)
	}

	if l.hadRuntimeError {
		return fmt.Errorf("%w. %w", errLox, err)
	}

	return nil
}

func (l *Lox) runPrompt() error {
	scanner := bufio.NewScanner(l.rw)

	_, err := l.rw.WriteString("> ")
	if err != nil {
		return fmt.Errorf("%w. %w", errLox, err)
	}

	err = l.rw.Flush()
	if err != nil {
		return fmt.Errorf("%w. %w", errLox, err)
	}

	for scanner.Scan() {
		err := l.run(scanner.Text())
		if err != nil {
			return fmt.Errorf("%w. %w", errLox, err)
		}

		l.hadError = false

		_, err = l.rw.WriteString("> ")
		if err != nil {
			return fmt.Errorf("%w. %w", errLox, err)
		}

		err = l.rw.Flush()

		if err != nil {
			return fmt.Errorf("%w. %w", errLox, err)
		}
	}

	return nil
}

// TODO: runtimeError - error -> RunTimeError.
func (l *Lox) runtimeError(errObj error) error {
	l.hadRuntimeError = true

	_, err := l.rw.WriteString(errObj.Error())
	if err != nil {
		return fmt.Errorf("%w. %w", errLox, err)
	}

	l.rw.Flush()

	err = l.rw.Flush()
	if err != nil {
		return fmt.Errorf("%w. %w", errLox, err)
	}

	return nil
}
