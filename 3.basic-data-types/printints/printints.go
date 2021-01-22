package main

import (
	"bytes"
	"fmt"
)

// intsToString is like fmt.Sprint([]int) but adds commas between each
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('<')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte('>')
	return buf.String()
}

func main() {
	v := intsToString([]int{4, 5, 2, 3, 1, 4, 9, 6})
	fmt.Println(v)
}
