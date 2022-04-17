package hitesh

import (
	"fmt"
	"sync"
)

func ChannelBasics() {
	wg := &sync.WaitGroup{}

	myChan := make(chan int, 2)

	wg.Add(1)
	go func() {
		fmt.Println(<-myChan)
		fmt.Println(<-myChan)
		wg.Done()
	}()

	go func() {
		myChan <- 2
		myChan <- 3
	}()

	wg.Wait()
}
