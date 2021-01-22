package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	s := `<!DOCTYPE html>
<html lang="en">
	<head>
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<title>Document</title>
	</head>
	<body>
		<h1>カタカナ</h1>
	</body>
</html>`

	sr := []rune(s)

	file, err := os.OpenFile("page.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Printf("%x\n", sr)
}
