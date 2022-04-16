package src

import "fmt"

func BufferChannelBasic() {
	mybufferchan := make(chan int, 2)
	mybufferchan <- 10
	fmt.Println(<-mybufferchan)
}
