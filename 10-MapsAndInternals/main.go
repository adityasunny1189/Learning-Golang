package main

import "fmt"

func main() {
	dict := map[string]string{
		"good": "bad",
		"can":  "cannont",
	}

	hindi := make(map[string]string)
	hindi["Hello"] = "namaskar"
	hindi["how are you"] = "kese hai aap"

	fmt.Println(hindi)
	fmt.Println(dict)
	fmt.Println(dict["good"])
	fmt.Println(len(dict))
}
