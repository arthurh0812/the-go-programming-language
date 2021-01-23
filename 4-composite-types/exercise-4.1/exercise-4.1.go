package main

import (
	"crypto/sha256"
	"fmt"
)

var popCount [256]byte

func init() {
	for i := range popCount {
		popCount[i] = popCount[i/2] + byte(i&1)
	}
}

func main() {
	sum1 := sha256.Sum256([]byte("x"))
	sum2 := sha256.Sum256([]byte("x"))

	fmt.Printf("%b\n", sum1)
	fmt.Printf("%b\n", sum2)

	n := checkSumBitDiff(sum1, sum2)

	fmt.Printf("%v\n", n)
}

func checkSumBitDiff(sum1, sum2 [32]byte) int {
	count := 0

	for i, b := range sum1 {
		count += int(popCount[b^sum2[i]])
	}

	return count
}
