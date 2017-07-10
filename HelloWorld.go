// HelloWorld
package main

import (
	"fmt"
)

func main() {
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
}
