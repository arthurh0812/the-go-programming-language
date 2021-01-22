package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

var wordfreq = make(map[string]int)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("no file name provided.")
	}

	file, err := os.OpenFile(os.Args[1], os.O_CREATE|os.O_RDWR, 666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := removePunctuation(file)

	input := bufio.NewScanner(buf)

	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		wordfreq[word]++
	}

	for word, freq := range wordfreq {
		fmt.Printf("%s\t%d\n", word, freq)
	}
}

func removePunctuation(file *os.File) *bytes.Buffer {
	byts, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var out = &bytes.Buffer{}

	for _, r := range string(byts) {
		if unicode.IsPunct(r) {
			out.WriteByte(32)
			continue
		}
		out.WriteRune(r)
	}

	return out
}
