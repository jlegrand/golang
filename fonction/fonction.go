package fonction

import (
	"math"
	"fmt"
)

func max(num1, num2 int) (result int) {

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return
}

func Max(num1, num2 int) int {
	return max(num1, num2)
}

func Swap(x, y string) (a, b string){
	a = y
	b = x
	return
}

func Anonymous() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}

