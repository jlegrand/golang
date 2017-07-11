package basics

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

var urls = []string {
	"http://golang.com",
	"http://google.com",
	"http://google.fr",
	"http://google.de",
}

func WaitGroup() {

	for _, url := range urls {

		// increment counter
		wg.Add(1)

		go func(url string) {
			// Decrement wait group
			defer wg.Done()
			fmt.Println(url)
		}(url)
	}

	wg.Wait()

}