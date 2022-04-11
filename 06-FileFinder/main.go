package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Provide a directory\n")
		return
	}
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	var totalSize int
	for _, file := range files {
		if file.Size() == 0 {
			totalSize += len(file.Name()) + 1
		}
	}

	names := make([]byte, 0, totalSize)
	fmt.Printf("total required size: %d\n", totalSize)

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
		}
	}
	err = ioutil.WriteFile("output.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(names)
}
