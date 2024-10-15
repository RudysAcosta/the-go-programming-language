package array

import (
	"crypto/sha256"
	"fmt"
)

func countDifferentBits(hash1, hash2 [32]byte) int {

	different := 0

	for i := 0; i < len(hash1); i++ {
		xor := hash1[i] ^ hash2[i]
		different += countBits(xor)
	}

	return different
}

func countBits(b byte) int {
	count := 0
	for b > 0 {
		count += int(b & 1)
		b >>= 1
	}
	return count
}

func CompareBit() {
	data1 := "Hello"
	data2 := "world"

	// calculate SHA256 hashes
	hash1 := sha256.Sum256([]byte(data1))
	hash2 := sha256.Sum256([]byte(data2))

	fmt.Println(hash1)
	fmt.Println(hash2)

	differentBits := countDifferentBits(hash1, hash2)

	fmt.Printf("Number of different bits: %d\n", differentBits)

}
