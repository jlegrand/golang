package basics

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

func Swap(x, y string) (a, b string) {
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

// déf un nouveau type
type Stringy func() string

func returnAFunction() Stringy {
	return func() string {
		fmt.Println("A l'interieur d'une Stringy fontion")
		return "bar"
	}
}

func TakesAFunction(foo Stringy) {
	fmt.Println("resultat de la fontion de type Stringy : ", foo())
}

func FonctionVariadique(is ...int) {
	for i := 0; i < len(is); i++ {
		fmt.Println(is[i])
	}
}
