// last page exercise:
// https://blog.golang.org/go-slices-usage-and-internals

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	filename := "text.txt"
	fmt.Println(copyDigits(filename))
	fmt.Println(appendDigits(filename))
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func copyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func appendDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	var c []byte
	return append(c, b...)
}
