package src

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func SetDS() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	rx := regexp.MustCompile(`[^a-z]+`)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")
		if len(word) > 2 {
			words[word] = true
		}
	}

	for word := range words {
		fmt.Println(word)
	}

	query := "venus"
	if words[query] {
		fmt.Printf("The input contains the query %s\n", query)
		return
	}
	fmt.Printf("The input does not contains the word\n")
}
