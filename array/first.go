package array

import (
	"crypto/sha256"
	"fmt"
)

func First() {

	var q [3]int = [3]int{1, 2}

	for i, v := range q {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
	fmt.Println()

	// Array  Literal
	r := [...]int{1, 2, 3}

	fmt.Println("<====Array Literal====>")

	fmt.Printf("Lenght Array: %d\n", len(r))

	for _, v := range r {
		fmt.Printf("Value: %d\n", v)
	}

	fmt.Println()

	//Compare array
	fmt.Println("<====Array Comparection====>")

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}

	fmt.Println(a == b, a == c, b == c)
	fmt.Println()

	// Sha256
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
