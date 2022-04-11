package src

import "fmt"

func TypeConversion() {
	speed := 100
	force := 2.5
	speed = int(float64(speed) * force)
	fmt.Println(speed)
}
