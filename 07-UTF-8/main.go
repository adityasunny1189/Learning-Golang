package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "πππππHello"

	bytes := []byte(str)

	bytes[0] = 'A'
	bytes[2] = 'B'

	fmt.Printf("size of %s id %d and rune size is %d\n",
		str, len(str), utf8.RuneCountInString(str))

	fmt.Printf("size of %s id %d and rune size is %d\n",
		bytes, len(bytes), utf8.RuneCount(bytes))
}
