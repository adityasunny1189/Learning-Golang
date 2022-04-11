package main

import (
	"bufio"
	"fmt"
	"log-parser/src"
	"os"
)

func InputScan() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	fmt.Printf("Scanned Text: %s", in.Text())

}

func main() {
	// InputScan()
	// src.SetDS()
	src.LogParser()
}
