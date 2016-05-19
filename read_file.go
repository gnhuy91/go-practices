// 1. Read a file line by line
// 2. Sort the lines
// ref:
// http://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
// http://stackoverflow.com/questions/8032170/how-to-assign-string-to-bytes-array

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	filename := "text.txt"
	_, strLines, err := readFileLine(filename)
	if err != nil {
		panic(err)
	}

	// remove blank lines from the slice
	for i, line := range strLines {
		if line == "\n" {
			strLines = append(strLines[:i], strLines[i+1:]...)
		}
	}
	// sort the slice by line
	sort.Strings(strLines)
	fmt.Println(strLines)
}

func readFileLine(filename string) ([]byte, []string, error) {
	var (
		byteLines []byte
		strLines  []string
	)

	// open the file
	o, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer o.Close()

	// start reading the file line by line
	r := bufio.NewReader(o)
	for {
		// ReadLine reads the line excluding newline characters
		l, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
		byteLines = append(byteLines, l...)
		// also append newline byte back so that result slice will include
		// newline characters like the original file
		byteLines = append(byteLines, []byte("\n")...)
		strLines = append(strLines, string(l)+"\n")
	}
	return byteLines, strLines, nil
}
