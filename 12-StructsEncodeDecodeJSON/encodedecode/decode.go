package encodedecode

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name       string          `json:"username"`
	Permission map[string]bool `json:"perms"`
}

func DecodeJSON() {
	var input []byte
	var users []user
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		input = append(input, in.Bytes()...)
	}
	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range users {
		fmt.Println(v)
	}
}
