package main

import "fmt"

func main() {
	s := "xello"

	boolean := s != "" && s[0] == 'x'
	fmt.Println(boolean)

	boolean = isASCII('4')
	fmt.Println(boolean)

	boolean = isASCII('*')
	fmt.Println(boolean)

	bit := boolBit(boolean)
	fmt.Println(bit)

	boolean = bitBool(uint8(bit))
	fmt.Println(boolean)
}

func isASCII(c rune) bool {
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
		return true
	}
	return false
}

func boolBit(boolean bool) int {
	if boolean {
		return 1
	}
	return 0
}

func bitBool(bit uint8) bool {
	return bit != 0
}
