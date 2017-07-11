// HelloWorld
package main

import (
	//"fmt"
	//"github.com/jlegrand/golang/boucles"
	//"github.com/jlegrand/golang/variables"
	//"github.com/jlegrand/golang/pointers"
	//"github.com/jlegrand/golang/fonction"
	"fmt"
)

func main() {

	//variables.Defvar()

	//boucles.Loop()
	//boucles.Switchcase()
	//boucles.Forloop()
	//boucles.Gotofunc()

	//fmt.Println(fonction.Max(10, 5))
	//fmt.Println(fonction.Swap("10", "5"))
	//fonction.Anonymous()
	//fmt.Println(fonction.Calcul('/', 12, 5))
	//fonction.Closure()

	/*fonction.FonctionVariadique(1)
	fonction.FonctionVariadique(1, 2)
	fonction.FonctionVariadique(1,2,3)
	// éclatement de la structure avec ...
	fonction.FonctionVariadique([]int{9,8,7,6,5,4,3,2,1}...)*/

	/*var t1 = []int{1,2,3}
	var t2 = []int{4,5,6}
	var t []int = append(t1,t2...)
	fmt.Println(t)*/

	//pointers.Swap(10, 5)

	//variables.Manipulate()
	//variables.Array()
	//variables.Slice()
	//variables.Append()

	fmt.Println(exercice1([]int {1,2,3,7,8,9}, 5))

}

func exercice1(array []int, value int) (result []int) {

	var i, l int = 0, len(array)
	for (i < l && array[i] < value) {
		i++
	}

	result = append(append(array[:i],value),array[i+1:]...)

	return
}


