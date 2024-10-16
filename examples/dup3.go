package examples

import (
	"fmt"
	"os"
	"strings"
)

func Dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			for _, world := range strings.Split(line, " ") {
				counts[world]++
			}
		}

		for world, n := range counts {
			fmt.Printf("%s => %d\n", world, n)
		}

	}

}
