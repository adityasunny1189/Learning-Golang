package src

import (
	"fmt"
	"unicode/utf8"
)

func StringLength() {
	name := "あでてぃや"
	len := utf8.RuneCountInString(name)
	fmt.Println(len)
}
