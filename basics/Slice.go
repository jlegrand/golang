package basics

import "fmt"

func Array() {

	var n [10]int

	for i := 0; i < 10; i++ {
		n[i]  = i + 100
	}

	for j,x := range n {
		fmt.Printf("Element[%d] = %d\n", j, x)
	}
}

func Slice() {
	var numbers = make([]int,3,5)
	printSlice(numbers)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}

func Append() {

	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	copy(numbers1, numbers)
	printSlice(numbers1)

}
