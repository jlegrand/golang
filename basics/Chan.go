package basics

import "fmt"

func Chan() {

	done := make(chan bool)

	go func() {
		fmt.Println("goroutine message")
		done <- true
	}()

	fmt.Println("main message")
	<-done
}
