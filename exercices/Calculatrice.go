package exercices

import (
	"fmt"
)

// Calculatrice tordue
type Operator func(int, int) int

//type error interface {
//	Error() string
//}

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

func OperationUsingMap(o rune) (fonction Operator, e error){

	var op map[rune]Operator
	op = make(map[rune]Operator, 4)

	op['+'] = func(a, b int) int { return a + b }
	op['-'] = func(a, b int) int { return a - b }
	op['*'] = func(a, b int) int { return a * b }
	op['/'] = func(a, b int) int { return a / b }

	fn, exists := op[o]
	if exists {
		return fn, nil
	} else {
		panic("Not implemented")
		//return nil, errors.New("Not Implemented")
	}

}

func Calcul(op rune, a, b int) {

	//return Operation(op)(a, b)
	o, error := OperationUsingMap(op)

	if(error == nil) {
		fmt.Println(o(a,b))
	} else {
		fmt.Println(error)
	}

}
