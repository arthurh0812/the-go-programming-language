package main

import "fmt"

func group(x []int) []int {
	z := make([]int, 0, len(x))
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		if i%2 == 0 {
			z = append(z, x[j], x[i])
		} else {
			z = append(z, x[i], x[j])
		}
	}
	return z
}

func main() {
	var sl []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(sl)

	for i, x := 1, sl; i <= len(sl); i++ {
		x = group(x)

		fmt.Printf("%v:\t%v\n", i, x)
	}

}
