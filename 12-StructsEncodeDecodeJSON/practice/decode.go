package practice

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func DecodeJSON() {
	fmt.Println("Decoding students details")
	var input []byte
	var students []Student
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		input = append(input, in.Bytes()...)
	}
	err := json.Unmarshal(input, &students)
	if err != nil {
		fmt.Println(err)
		return
	}
	queryStudent := "Amit Shahwal"
	ok := false
	for _, student := range students {
		if student.Name == queryStudent {
			ok = true
			for _, placementOffer := range student.Offers {
				if placementOffer.Fte {
					fmt.Printf("got full time offer of %f from %s\n",
						placementOffer.Package, placementOffer.CompanyName)
				}
			}
		}
	}
	if !ok {
		fmt.Printf("student details not present\n")
	}
}
