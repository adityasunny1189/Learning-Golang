package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	MAXTURNS = 5
)

func luckyNumber() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 2 {
		fmt.Printf("pick a number\n")
		return
	}
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("error in input\n")
		return
	}
	if guess < 0 {
		fmt.Printf("input must be greater than 0\n")
		return
	}
	for turn := 0; turn < MAXTURNS; turn++ {
		i := rand.Intn(guess + 1)
		if i == guess && turn == 0 {
			fmt.Printf("YOU WIN IN FIRST CHANCE\n")
			return
		} else if i == guess {
			fmt.Printf("YOU WIN\n")
			return
		}
	}
	fmt.Printf("YOU LOSE\n")
}

func main() {
	fmt.Printf("Project on GO Basics\n")
	luckyNumber()
}
