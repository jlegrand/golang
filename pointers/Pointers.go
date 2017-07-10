package pointers

import "fmt"

func DemoPointers() {

	var v int = 1234
	//fmt.Println("pointeur de v : ", *v)
	fmt.Println("valeur de v : ", v)
	fmt.Println("adresse de v : ", &v)

	var v2 *int = &v
	fmt.Println("pointeur de v2 : ", *v2)
	fmt.Println("valeur de v2 : ", v2)
	fmt.Println("adresse de v2 : ", &v2)

	var v3 **int = &v2
	fmt.Println("pointeur de v3 : ", *v3)
	fmt.Println("valeur de v3 : ", v3)
	fmt.Println("adresse de v3 : ", &v3)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x  // save the value of x
	*x = *y    // put y into x
	*y = temp  // put temp into y
}

func Swap(a, b int) {
	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	swap(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
}
