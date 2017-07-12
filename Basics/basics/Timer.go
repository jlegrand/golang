package basics

import (
	"time"
	"fmt"
	"math/rand"
)

func Timer() {
	start := time.Now()
	var t *time.Timer

	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println((time.Now().Sub(start)))
		t.Reset(randomDuration())
	})
	time.Sleep(2 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(10))
}
