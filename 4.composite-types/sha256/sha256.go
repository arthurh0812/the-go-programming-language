package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("helloworld"))
	fmt.Printf("%x\n%x\n%t\n%[1]T\n%[2]T\n", c1, c2, c1 == c2)

}
