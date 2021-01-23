package main

import (
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := os.Args[1]
	l := utf8.RuneCountInString(msg)
	s := strings.Repeat("!", l)
	result := s + msg + s
	println(result)

}
