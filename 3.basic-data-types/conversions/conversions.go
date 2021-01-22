package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // 123 123

	fmt.Println(strconv.FormatInt(int64(x), 2)) // 1111011

	fmt.Printf("x=%b\n", x)

	z, err := strconv.Atoi("234")
	if err != nil {
		log.Fatal(err)
	}
	a, err := strconv.ParseInt("234", 10, 64) // base 10, up to 64 bits
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(z, a)
}
