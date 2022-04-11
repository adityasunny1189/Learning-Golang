package pkg

import "fmt"

const (
	summer = 3
	winter = 1
	yearly = summer + winter
)

func Books() {
	var books [yearly]string
	var sBooks [summer]string
	var wBooks [winter]string

	books[0] = "Atomic Habits"
	books[1] = "Eat that Frog"
	books[2] = "5AM Club"
	books[3] = "India's most fearless"

	wBooks[0] = books[0]

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("%#v\n", books)
	fmt.Printf("%#v\n", sBooks)
	fmt.Printf("%#v\n", wBooks)
}
