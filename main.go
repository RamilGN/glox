package main

import (
	"log"
	"os"
)

func main() {
	switch len(os.Args) {
	case 2:
		err := NewLox(nil).Start([]string{os.Args[1]})
		if err != nil {
			log.Fatal(err)
		}
	case 1:
		err := NewLox(nil).Start([]string{})
		if err != nil {
			log.Fatal(err)
		}
	}
}
