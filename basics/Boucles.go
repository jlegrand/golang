package basics

import "fmt"

func Loop() {
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

func Switchcase() {

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

func Forloop() {

	var a int
	var b int = 15
	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("value of x = %d at i = %d\n", x, i)
	}
}

func Gotofunc() {

	var a int = 10

LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}

		fmt.Printf("value of a : %d\n", a)
		a++
	}

}
