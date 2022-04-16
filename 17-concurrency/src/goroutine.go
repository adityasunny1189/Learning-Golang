package src

import (
	"fmt"
	"time"
)

func HelloWorld() {
	fmt.Println("welcome to go routine")
}

func GoRoutineBasics() {
	go HelloWorld()
	go func() {
		fmt.Println("hello world")
	}()

	time.Sleep(time.Millisecond * 200)
}
