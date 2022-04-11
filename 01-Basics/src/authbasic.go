package src

import (
	"fmt"
	"os"
)

func CheckAuth() {
	if len(os.Args) == 3 {
		Authentication()
	} else {
		AccessDenied()
	}
}

func Authentication() {
	u, p := "adityapathak", "pass1234"
	username, password := os.Args[1], os.Args[2]
	if username == u {
		if password == p {
			fmt.Printf("User Authenticated\n")
		} else {
			fmt.Printf("Incorrect Password for %s\n", username)
		}
	} else {
		fmt.Printf("Access denied for %s\n", username)
	}
}

func AccessDenied() {
	fmt.Printf("Usage: [username] [password]\n")
}
