package src

import "fmt"

func DaysOfWeek() {
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}

func TimeZone() {
	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)
	fmt.Println(EST, MST, PST)
}
