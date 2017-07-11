package basics

import "fmt"

func Defvar() {

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
