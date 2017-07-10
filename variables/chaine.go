package variables

import "fmt"

func ManipulateString() {

	var greeting = "Hello World !"

	fmt.Printf("%s", greeting)
	fmt.Printf("\n")
	fmt.Printf("hex bytes: ")
	for i:=0; i < len(greeting); i++ {
		fmt.Printf("%x", greeting[i])
	}
	fmt.Printf("\n")

	const sampleText = "\xdb\xb2\x3d\x20\xe2\x8c\x98"
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")

}
