package main

import (
	"fmt"

	basename "github.com/arthurh0812/the-go-programming-language/3-basic-data-types/basename/basename"
)

func main() {
	b1 := basename.Basename1("files/packages/anything/okandnotok.rust")
	fmt.Println(b1) // "okandnotok"
	b2 := basename.Basename2("morefiles/alright.go")
	fmt.Println(b2) // "alright"
}
