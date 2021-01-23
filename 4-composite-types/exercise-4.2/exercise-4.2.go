package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var selectedHash string

	flag.StringVar(&selectedHash, "hash", "sha256", "hashing algorithm")

	flag.Parse()

	args := flag.Args()
	var input []byte
	if len(args) > 0 {
		input = []byte(args[0])
	}

	switch selectedHash {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256(input))
		return
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384(input))
		return
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(input))
	default:
		fmt.Println("Unknown hash...")
		return
	}
}
