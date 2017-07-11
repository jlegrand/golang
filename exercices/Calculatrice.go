package exercices

import "fmt"

// Calculatrice tordue
type Operator func(int, int) int

func Operation(op rune) (fonction Operator) {

	switch op {
	case '+':
		fonction = func(a, b int) int { return a + b }
	case '-':
		fonction = func(a, b int) int { return a - b }
	case '*':
		fonction = func(a, b int) int { return a * b }
	case '/':
		fonction = func(a, b int) int { return a / b }
	default:
		fmt.Println("Not implemented")
	}

	return
}

func Calcul(op rune, a, b int) int {

	return Operation(op)(a, b)

}
