package basics

import (
	"fmt"
	"time"
)

func Routine() {
	go fmt.Println("World !")
	fmt.Println("Hello ")
	time.Sleep((10 * time.Microsecond))
}

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}()
}
