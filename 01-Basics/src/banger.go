package src

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func Banger() {
	msg := os.Args[1]
	l := utf8.RuneCountInString(msg)
	s := strings.Repeat("!", l)
	ns := s + strings.ToUpper(msg) + s
	fmt.Println(ns)
}
