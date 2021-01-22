package main

import "fmt"

func main() {
	cs := comma("3925")

	fmt.Println(cs)
}

// comma inserts commas in a non-negative decimal interger string.
func comma(s string) string {
	length := len(s)
	if length <= 3 {
		return s
	}
	return comma(s[:length-3]) + "," + s[length-3:]
}
