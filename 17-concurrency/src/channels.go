package src

import (
	"fmt"
)

func SendData(mc chan int) {
	for i := 0; i < 5; i++ {
		mc <- i
	}
	close(mc)
}

func ChannelBasics() {
	myChannel := make(chan int)
	defer close(myChannel)
	go SendData(myChannel)
	// for i := 0; i < 6; i++ {
	// 	val, open := <-myChannel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(val)
	// }
	for val := range myChannel {
		fmt.Println(val)
	}
}
