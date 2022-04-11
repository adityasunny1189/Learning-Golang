package src

import "fmt"

func CountryName() {
	city := "Mumbai"
	var country string
	switch city {
	case "Paris":
		country = "France"
	case "London":
		country = "England"
	case "Delhi", "Mumbai":
		country = "India"
	default:
		country = "Country not found"
	}
	fmt.Printf("country: %s\n", country)
}

func NumbersSwitch() {
	i := 142
	switch {
	case i > 100:
		fmt.Printf("greater than 100\n")
		fallthrough
	case i > 200:
		fmt.Printf("greater than 200\n")
		fallthrough
	default:
		fmt.Printf("no is no")
	}
}
