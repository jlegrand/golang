// HelloWorld
package main

import (
	"fmt"
)

func main() {

	//defvar()
	//loop()
	switchcase()

}

func defvar() {

	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T \n", x)
	fmt.Printf("x value is %f \n", x)

	var chaine string
	chaine = "ma chaine en utf-8"
	fmt.Println(chaine)
	fmt.Printf("%s \n", chaine)

	var char rune
	char = 'r'
	fmt.Printf("char is of type %T \n", char)
	fmt.Printf("%d \n", char)
	fmt.Printf("%q \n", char)

	e := 9000.7
	fmt.Printf("e is of type %T \n", e)

	var a, b, c = 1, 8.0, "etet"
	fmt.Printf("a is of type %T \n", a)
	fmt.Printf("b is of type %T \n", b)
	fmt.Printf("c is of type %T \n", c)

	const LENGTH int = 10

}

func loop() {
	var a int = 10

	if a < 20 {
		fmt.Printf("a is less than 20\n")
	}

	if limit := 8; a < limit {
		fmt.Printf("a is less than %d\n", limit)
	} else if limit = 5; a > limit {
		fmt.Printf("a is greater than %d\n", limit)
	} else {
		fmt.Printf("a is less than %d\n", limit)
	}
}

func switchcase() {

	var grade string = "B"
	var marks int = 80

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("Excellent")
	case grade == "B", grade == "C":
		fmt.Println("Well done")
	default:
		fmt.Println("Invalid grade")
	}
}
