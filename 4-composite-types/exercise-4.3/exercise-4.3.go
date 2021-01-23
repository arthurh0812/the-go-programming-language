package main

import "fmt"

func reverse(ar *[10]int) {
	for i, j := 0, 9; i < j; i, j = i+1, j-1 {
		ar[i], ar[j] = ar[j], ar[i]
	}
}

func main() {
	ar := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	reverse(&ar)

	fmt.Println(ar)
}
