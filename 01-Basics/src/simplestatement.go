package src

import (
	"fmt"
	"os"
	"strconv"
)

func SimpleStatementFunc() {
	if _, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("can not convert\n")
	} else {
		fmt.Printf("yeah\n")
	}
}
