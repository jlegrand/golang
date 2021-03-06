package basics

import "fmt"

func Defer() {
	f()
	fmt.Println("return normally from f")
}

func f() {
	defer func() {
		if r:= recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("calling g")
	g(0)
	fmt.Println("return normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g ", i)
	fmt.Println("Printing in g ", i)
	g(i+1)
}

func OrderedDefer() {

	defer func() {
		fmt.Println("defer 1")
	}()

	defer func() {
		fmt.Println("defer 2")
	}()

	defer func() {
		fmt.Println("defer 3")
	}()
}
