// HelloWorld
package main

import (
	//"fmt"
	//"github.com/jlegrand/golang/boucles"
	//"github.com/jlegrand/golang/variables"
	//"github.com/jlegrand/golang/pointers"
	"github.com/jlegrand/golang/fonction"
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
	fonction.FonctionVariadique(1)
	fonction.FonctionVariadique(1, 2)
	fonction.FonctionVariadique(1,2,3)
	// éclatement de la structure avec ...
	fonction.FonctionVariadique([]int{9,8,7,6,5,4,3,2,1}...)

	//pointers.Swap(10, 5)

	//variables.Manipulate()
	//variables.Array()
	//variables.Slice()
	//variables.Append()

}


