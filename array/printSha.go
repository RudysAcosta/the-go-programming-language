package array

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func PrintSha() {

	useSHA384 := flag.Bool("sha384", false, "Use SHA384 hashing algorithm")
	useSHA512 := flag.Bool("sha512", false, "Use Sha512 hashing algorithm")

	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	switch {
	case *useSHA384:
		hash := sha512.Sum384(data)
		fmt.Printf("SHA384: %x\n", hash)
	case *useSHA512:
		hash := sha512.Sum512(data)
		fmt.Printf("SHA512: %x\n", hash)
	default:
		hash := sha256.Sum256(data)
		fmt.Printf("SHA256: %x\n", hash)
	}
}
