package src

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func UniqueRandomNumber() {
	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])

	var uniques []int

loop:
	for len(uniques) < max {
		n := rand.Intn(max)
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques = append(uniques, n)
	}
	sort.Ints(uniques)
	fmt.Println("\n\n", uniques)
}
