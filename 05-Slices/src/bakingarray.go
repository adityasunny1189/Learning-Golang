package src

import "fmt"

func BakingArray() {
	ages := []int{12, 21, 34}
	age1 := ages[:1]
	age2 := ages[1:2]
	fmt.Println(ages, age1, age2)
	ages[0] = 123
	fmt.Println(ages, age1, age2)
	fmt.Println(age1[1])
}
