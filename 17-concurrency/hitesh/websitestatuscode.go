package hitesh

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals []string

func StatusCodeFunc() {
	websiteslist := []string{
		"https://google.com",
		"https://github.com",
		"https://fb.com",
		"https://go.dv",
	}

	for _, website := range websiteslist {
		go GetStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func GetStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Website Down")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("status code of %s is %d\n", endpoint, res.StatusCode)
	}
	wg.Done()
}
