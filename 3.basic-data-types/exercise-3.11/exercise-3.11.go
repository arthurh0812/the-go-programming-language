package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	formattedNum := comma(5645768626.6700001)

	fmt.Println(formattedNum)
}

// comma formats the given number with added commas.
func comma(num float64) string {
	str := strconv.FormatFloat(num, 'f', -1, 64)

	length := len(str)

	dot := strings.LastIndexByte(str, '.')

	var buf bytes.Buffer

	for i, j := dot-1, 0; i >= 0 && j < length; {
		if j != 0 && (i+1)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(str[j])
		i--
		j++
	}

	buf.WriteString(str[dot:])

	return buf.String()
}
