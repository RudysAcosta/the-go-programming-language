package basic

import (
	"fmt"
	"os"
)

func Command() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " \n"
	}

	fmt.Println(s)
}
