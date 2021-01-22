package main

import (
	"fmt"
)

func remove(arr []int, index int) []int {
	copy(arr[index:], arr[index+1:])
	return arr[:len(arr)-1]
}

func remove2(arr []int, index int) []int {
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 3))

	t := []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(t, 1))
}
