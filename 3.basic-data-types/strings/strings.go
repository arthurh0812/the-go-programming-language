package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	s := "Hello, World!"
	fmt.Println(len(s))
	fmt.Println(s[0], s[len(s)-1])
	// fmt.Println(s[len(s)])

	// substring
	fmt.Println(s[0:5])
	fmt.Println(s[:])
	fmt.Println("Goodbye" + s[5:])

	s = "界"
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])

	html := `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=, initial-scale=1.0" />
		<title>Strings</title>
		<link rel="stylesheet" src="./mystyle.css">
	</head>
	<body>
		<h1>Hello, World 界!</h1>
	</body>
</html>`

	fname := "page.html"
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	file.Write([]byte(html))

}
