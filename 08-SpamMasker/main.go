package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Provide input\n")
		return
	}
	text := args[0]

	const (
		link     = "http://"
		linkSize = len(link)
	)

	buf := make([]byte, 0, len(text))
	endIndex := 0

	for i := 0; i < len(text); i++ {
		buf = append(buf, text[i])
		if len(text[i:]) > linkSize && text[i:i+linkSize] == link {
			endIndex = i + linkSize
		}
	}
	for i := endIndex; ; i++ {
		if buf[i] == ' ' || buf[i] == '\n' {
			break
		}
		buf[i] = '*'
	}
	fmt.Printf("Buffer is %s\n", buf)
}
