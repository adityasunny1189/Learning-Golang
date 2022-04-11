package src

import (
	"fmt"
	"os"
	"strconv"
)

func Loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("loop running %d time\n", i)
	}
}

func PrintEven() {
	n := 1
	for {
		if n > 100 {
			break
		}
		if n%2 != 0 {
			n++
			continue
		}
		fmt.Printf("%d\t", n)
		n++
	}
}

func MultiplicationTable() {
	n, _ := strconv.Atoi(os.Args[1])
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Printf("\n")
	}
}

func ForRangeLoop() {
	for i, v := range os.Args[1:] {
		fmt.Printf("%d\t%s\n", i, v)
	}
}
