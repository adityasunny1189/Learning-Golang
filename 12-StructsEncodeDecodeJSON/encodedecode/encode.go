package encodedecode

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type User struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func EncodeJSON() {
	users := []User{
		{"aditya", "1234", nil},
		{"sunny", "4erw", permissions{"admin": true}},
		{"kunal", "231w", permissions{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
