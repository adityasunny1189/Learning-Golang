package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	summer = 3
	winter = 1
	yearly = summer + winter
)

func Books() {
	books := [...]string{
		"Atomic Habits",
		"Eat that Frog",
		"5AM Club",
		"India's most fearless",
	}
	var sBooks [summer]string
	var wBooks [winter]string

	// books[0] = "Atomic Habits"
	// books[1] = "Eat that Frog"
	// books[2] = "5AM Club"
	// books[3] = "India's most fearless"

	wBooks[0] = books[0]

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("%#v\n", books)
	fmt.Printf("%#v\n", sBooks)
	fmt.Printf("%#v\n", wBooks)
}

func RandomMood() {
	if len(os.Args[1:]) != 1 {
		fmt.Printf("usage: [name]\n")
		return
	}

	name := os.Args[1]

	// moods := [...]string{
	// 	"good", "happy", "awesome",
	// 	"bad", "sad", "terrible",
	// }

	moods := [...][3]string{
		{"good", "happy", "awesome"},
		{"bad", "sad", "terrible"},
	}

	rand.Seed(time.Now().UnixNano())
	mood := moods[rand.Intn(len(moods))][rand.Intn(len(moods[0]))]
	fmt.Printf("%s is feeling %s\n", name, mood)
}

func MultiDimensionalArray() {
	// arr := [2][3]int{
	// 	[3]int{1, 2, 3},
	// 	[3]int{4, 3, 2},
	// }

	arr := [...][3]int{
		{1, 2, 3},
		{4, 3, 1},
	}

	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Printf("\n")
	}
}

func KeyedElement() {
	rates := [...]float64{
		0: 0.5,
		1: 1.5,
		2: 6.2,
	}
	fmt.Printf("%v", rates)
}

func main() {
	fmt.Printf("Array Basics\n")
	RandomMood()
	// MultiDimensionalArray()
}
