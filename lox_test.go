package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	l := NewLox(nil)

	assert.IsType(t, &Lox{}, l)
	assert.NotNil(t, l.rw)
}

func TestStart(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		rb := &bytes.Buffer{}
		wb := &bytes.Buffer{}
		rb.WriteString("foo")
		rw := bufio.NewReadWriter(bufio.NewReader(rb), bufio.NewWriter(wb))
		l := NewLox(rw)
		err := l.Start([]string{})
		assert.Nil(t, err)
		assert.Equal(t, "> foo\n> ", wb.String())
	})

	t.Run("one arg(filename)", func(t *testing.T) {
		file, err := os.CreateTemp("", "start_test")
		require.NoError(t, err)
		defer file.Close()

		_, err = file.WriteString("foo")
		require.NoError(t, err)

		rb := &bytes.Buffer{}
		wb := &bytes.Buffer{}
		rw := bufio.NewReadWriter(bufio.NewReader(rb), bufio.NewWriter(wb))
		l := NewLox(rw)

		err = l.Start([]string{file.Name()})
		assert.Nil(t, err)
		assert.Equal(t, "foo\n", wb.String())
	})

	t.Run("too many args", func(t *testing.T) {
		rw := bufio.NewReadWriter(bufio.NewReader(&bytes.Buffer{}), bufio.NewWriter(&bytes.Buffer{}))
		l := NewLox(rw)
		err := l.Start([]string{"foo", "bar"})
		assert.NotNil(t, err)
	})
}
