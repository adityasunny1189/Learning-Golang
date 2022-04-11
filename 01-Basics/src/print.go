package src

import "fmt"

func PrintFunc() {
	var (
		planet   = "Earth"
		distance = 261
		orbital  = 321.34675
		hasLife  = true
	)

	fmt.Printf("%s is %d km away and has orbital %.3f and life present is %t",
		planet, distance, orbital, hasLife)
}
