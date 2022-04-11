package src

import (
	"fmt"
	"os"
)

func ReadCommandLineArgs() {
	fmt.Println(os.Args[0])
}

func Greeter() {
	fmt.Println("hello ", os.Args[1])
}
