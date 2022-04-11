package src

import "fmt"

func SliceBasics() {
	var books []string
	books = append(books, "dracula")
	fmt.Printf("%s\n", books[0])
}
