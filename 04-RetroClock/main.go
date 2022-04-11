package main

import (
	"fmt"
	"time"
)

func main() {
	type placeholder [5]string
	zero := placeholder{
		"📦📦📦",
		"📦  📦",
		"📦  📦",
		"📦  📦",
		"📦📦📦",
	}
	one := placeholder{
		"📦📦 ",
		"  📦 ",
		"  📦 ",
		"  📦 ",
		"📦📦📦",
	}

	two := placeholder{
		"📦📦📦",
		"    📦",
		"  📦📦",
		"📦    ",
		"📦📦📦",
	}

	three := placeholder{
		"📦📦📦",
		"    📦",
		"📦📦📦",
		"    📦",
		"📦📦📦",
	}

	four := placeholder{
		"📦  📦",
		"📦  📦",
		"📦📦📦",
		"    📦",
		"    📦",
	}

	five := placeholder{
		"📦📦📦",
		"📦   ",
		"📦📦📦",
		"    📦",
		"📦📦📦",
	}

	six := placeholder{
		"📦📦📦",
		"📦   ",
		"📦📦📦",
		"📦  📦",
		"📦📦📦",
	}

	seven := placeholder{
		"📦📦📦",
		"    📦",
		"    📦",
		"    📦",
		"    📦",
	}

	eight := placeholder{
		"📦📦📦",
		"📦  📦",
		"📦📦📦",
		"📦  📦",
		"📦📦📦",
	}

	nine := placeholder{
		"📦📦📦",
		"📦  📦",
		"📦📦📦",
		"    📦",
		"📦📦📦",
	}

	colon := placeholder{
		"    ",
		" 📦 ",
		"    ",
		" 📦 ",
		"    ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for {
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[min%10],
		}

		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}

}
