package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Lox struct {
	rw *bufio.ReadWriter
}

func NewLox(rw *bufio.ReadWriter) *Lox {
	if rw == nil {
		rw = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	}
	l := Lox{rw: rw}
	return &l
}

func (l *Lox) Start(args []string) error {
	switch len(args) {
	case 1:
		err := l.runFile(args[0])
		if err != nil {
			return err
		}
	case 0:
		err := l.runPrompt()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Usage: glox [filename]")
	}

	return nil
}

func (l *Lox) run(src string) error {
	scanner := bufio.NewScanner(strings.NewReader(src))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		_, err := l.rw.WriteString(scanner.Text())
		if err != nil {
			return err
		}
		_, err = l.rw.WriteString("\n")
		if err != nil {
			return err
		}

		err = l.rw.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lox) runFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = l.run(string(content))
	if err != nil {
		return err
	}

	return nil
}

func (l *Lox) runPrompt() error {
	scanner := bufio.NewScanner(l.rw)
	_, err := l.rw.WriteString("> ")
	if err != nil {
		return err
	}
	err = l.rw.Flush()
	if err != nil {
		return err
	}

	for scanner.Scan() {
		err := l.run(scanner.Text())
		if err != nil {
			return err
		}
		_, err = l.rw.WriteString("> ")
		if err != nil {
			return err
		}
		err = l.rw.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}
