package hitesh

import (
	"fmt"
	"sync"
)

func RaceConditionLearning() {
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	nums := []int{1, 2, 3}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("1st goroutine")
		mut.Lock()
		nums = append(nums, 11)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("2nd goroutine")
		mut.Lock()
		nums = append(nums, 22)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("3rd goroutine")
		mut.Lock()
		nums = append(nums, 33)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("3rd goroutine")
		mut.RLock()
		fmt.Println(nums)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(nums)
}
