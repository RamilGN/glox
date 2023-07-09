package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	l := NewLox(nil)

	assertType(t, &Lox{}, l)
	assertNotNil(t, l.rw)
}

func TestStart(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		rb := &bytes.Buffer{}
		wb := &bytes.Buffer{}
		rb.WriteString("foo")
		rw := bufio.NewReadWriter(bufio.NewReader(rb), bufio.NewWriter(wb))
		l := NewLox(rw)
		err := l.Start([]string{})

		assertNil(t, err)
		assertEqualString(t, "> identifier foo <nil>\neof  <nil>\n> ", wb.String())
	})

	t.Run("one arg(filename)", func(t *testing.T) {
		file, err := os.CreateTemp("", "start_test")
		requireNil(t, err)
		defer file.Close()

		_, err = file.WriteString("foo")
		requireNil(t, err)

		rb := &bytes.Buffer{}
		wb := &bytes.Buffer{}
		rw := bufio.NewReadWriter(bufio.NewReader(rb), bufio.NewWriter(wb))
		l := NewLox(rw)
		err = l.Start([]string{file.Name()})

		assertNil(t, err)
		assertEqualString(t, "identifier foo <nil>\neof  <nil>\n", wb.String())
	})

	t.Run("too many args", func(t *testing.T) {
		rw := bufio.NewReadWriter(bufio.NewReader(&bytes.Buffer{}), bufio.NewWriter(&bytes.Buffer{}))
		l := NewLox(rw)
		err := l.Start([]string{"foo", "bar"})

		assertNotNil(t, err)
	})
}
