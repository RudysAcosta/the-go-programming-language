package array

import "fmt"

func Run() {

	q := []int{1, 2, 3, 4, 5, 6}
	// reverseExercise(&q)
	// fmt.Println("Exercise 4.3")
	// fmt.Println(q)

	// // newIndex := (len(q) - 1) - (0 - 2)

	// fmt.Println(len(q))

	// rotateExercise(q, 2)

	fmt.Println(q)

	fmt.Println(6 % 6)
}

// Rewrite reverse to use an array pointer instead of a slice
func reverseExercise(p []int) {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}

func rotateToLeft(arr []int, k int) {
	n := len(arr)
	k = k % n

	fmt.Printf("n: %d, k: %d\n\n", n, k)

	reverseExercise(arr)
	reverseExercise(arr[:k])
	reverseExercise(arr[k:])
}

func rotalToRight(nums []int, k int) {

	n := len(nums)
	k = k % n

	reverseExercise(nums)
	reverseExercise(nums[:k])
	reverseExercise(nums[:k])
}
