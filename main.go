package main

import (
	"log"
	"os"
)

const (
	withPrompt = iota + 1
	withFile
)

func main() {
	switch len(os.Args) {
	case withPrompt:
		err := NewLox(nil).Start([]string{})
		if err != nil {
			log.Fatal(err)
		}
	case withFile:
		err := NewLox(nil).Start([]string{os.Args[1]})
		if err != nil {
			log.Fatal(err)
		}
	}
}
