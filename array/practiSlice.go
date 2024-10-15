package array

import "fmt"

func PractiSlice() {

	months := [...]string{
		"", "January", "February", "March", "April", "May",
		"June", "July", "August", "September", "October",
		"November", "December",
	}

	q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(q2)
	fmt.Println(summer)

	fmt.Println(summer[:20])

	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)

	// Reverse reverces a slice of ints in place
	fmt.Println()

}

func reverse(s []int) []int {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
