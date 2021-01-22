package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	// cwd, err := os.Getwd() compile error: unused cwd
	// if err != nil {
	// 	log.Fatal("os.Getwd failed: %v", err)
	// }

	// cwd, err := os.Getwd() Doesn't
	// if err != nil {
	// 	log.Fatalf("os.Getwd failed: %v", err)
	// }
	// log.Printf("Working directory = %s", cwd) !just a trick!

	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

func f() {}

var g = "g"

var fname string = "nothing.exe"

func main() {
	f := "f"
	fmt.Println(f)
	fmt.Println(g)
	// fmt.Println(h) compile error: undefined h

	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}

	y := "hello"
	for _, y := range y {
		y := y + 'A' - 'a'
		fmt.Printf("%c", y)
	}

	{
		z := "ok"
		fmt.Println(z)
	}

	// fmt.Println(z) compile error: undefined z

	ifs()

	fileBad()

	fileGood1()

	fileGood2()
}

func ifs() {
	if x := h(); x == 0 {
		fmt.Println(x)
	} else if y := j(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}

	// fmt.Println(x, y) compile error: x and y are not visible here
}

func h() int {
	return 10
}

func j(y int) int {
	return y * 1
}

func fileBad() error {
	// if f, err := os.Open(fname); err != nil {
	// 	// compile error: unused f
	// 	return err
	// }
	// f.Stat() compile error: undefined f
	// f.Close() compile error: undfined f
	return nil
}

func fileGood1() error {
	f, err := os.Open(fname)
	if err != nil {
		return nil
	}
	f.Stat()
	f.Close()
	return nil
}

func fileGood2() error {
	if f, err := os.Open(fname); err != nil {
		return err
	} else {
		f.Stat()
		f.Close()
		return nil
	}
}
