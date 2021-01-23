package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	formattedNumber := comma(32456)
	fmt.Println(formattedNumber)
}

func comma(num int) string {
	str := strconv.Itoa(num)

	length := len(str)
	if length <= 3 {
		return str
	}

	var buf bytes.Buffer

	for i, j := length-1, 0; i >= 0 && j < length; {
		if j != 0 && (i+1)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(str[j])
		i--
		j++
	}

	return buf.String()
}
