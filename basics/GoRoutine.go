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
