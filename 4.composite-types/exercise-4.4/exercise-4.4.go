package main

import "fmt"

func reverse(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func rotateRight(nums []int, n int) {
	// first reverse all
	reverse(nums)
	reverse(nums[n:])
	reverse(nums[:n])
}

func rotateLeft(nums []int, n int) {
	reverse(nums[n:])
	reverse(nums[:n])
	// in the end reverse all
	reverse(nums)
}

func main() {
	items := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	rotateRight(items, 4)

	fmt.Println(items)

	rotateLeft(items, 5)

	fmt.Println(items)
}
