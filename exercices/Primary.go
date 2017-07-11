package exercices

import (
	"fmt"
)

func Primary() {

	c := search(2, 100)
	for e := range c {
		fmt.Println(e)
	}
}

func Primary2() {

	in := make(chan int)
	out := make(chan int)
	
	primaryWithRoutine(in, out)

	in <- 3
	in <- 100

	for e := range out {
		fmt.Println(e)
	}
}

func primaryWithRoutine(in chan int, out chan int) () {

	go func() {
		min := <- in
		max := <- in
		c := search(min, max)
		for e := range c {
			out <- e
		}
		close(out)
	}()
}

func search(min, max int) <-chan int {

	c := make(chan int)

	go func() {
		for i := min; i < max; i++ {
			if isPrime(i) {
				c <- i
			}
		}
		close(c)
	}()
	return c
}

func isPrime(a int) (result bool) {
	result = true
	for i := 2; i <= a/2; i++ {
		if a % i == 0 {
			result = false
		}
	}
	return
}
