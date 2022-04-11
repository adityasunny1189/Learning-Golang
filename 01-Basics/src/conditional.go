package src

import "fmt"

func Conditionals() {
	score, valid := 5, true
	if score > 3 && valid {
		fmt.Printf("good\n")
	} else if score == 3 && !valid {
		fmt.Printf("oops\n")
	} else {
		fmt.Printf("not good\n")
	}
}
