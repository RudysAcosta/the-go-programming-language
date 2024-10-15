package exercises

import (
	"bufio"
	"fmt"
	"os"
)

func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++

		if len(counts) > 3 {
			break
		}
	}

	fmt.Println(counts)

	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d \t%s\n", n, line)
	// 	}
	// }
}
