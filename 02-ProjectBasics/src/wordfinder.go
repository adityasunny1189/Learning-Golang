package src

import (
	"fmt"
	"os"
	"strings"
)

const (
	wordList = "Hello world this is aditya"
)

func WordFinder() {
	words := strings.Fields(wordList)
	query := os.Args[1:]
queries:
	for _, q := range query {
	search:
		for i, w := range words {
			switch w {
			case "world", "this":
				break search
			}
			if w == strings.ToLower(q) {
				fmt.Printf("#%-2d: %s\n", i+1, q)
				break queries
			}
		}
	}
}
