package exercices

import "fmt"

func SliceInsertion(array []int, value int) (result []int) {

	var i, l int = 0, len(array)
	for (i < l && array[i] < value) {
		i++
	}

	fmt.Println(i)

	var head, tail []int = array[:i], array[i:]
	result = make([]int, i)

	fmt.Println(head)
	fmt.Println(tail)
	copy(result, head)
	fmt.Println(result)
	result = append(result, value)
	fmt.Println(result)
	result = append(result, tail...)
	fmt.Println(result)

	return
}
