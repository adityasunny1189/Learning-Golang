package src

import (
	"fmt"
	"time"
)

func GreetWishes() {
	switch ch := time.Now().Hour(); {
	case ch >= 18:
		fmt.Printf("good evening")
	case ch >= 12:
		fmt.Printf("good afternoon")
	case ch >= 6:
		fmt.Printf("good morning")
	default:
		fmt.Printf("good night")
	}
}
