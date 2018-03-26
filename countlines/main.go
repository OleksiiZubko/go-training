package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

// CountLines count lines of content in the file
func CountLines(path string) (int, error) {
	r, err := os.Open(path)

	if err != nil {
		println("Can't open file.")
		println(err.Error())

		defer r.Close()

		return 0, err
	}

	lines, err := lineCounter(r)

	if err != nil {
		println("Can't count files")
		println(err.Error())
		return 0, err
	}

	return lines, nil
}

func main() {
	count, err := CountLines("testdata/alice.txt")

	if err != nil {
		println("Error: ")
		println(err.Error())
	}

	fmt.Printf("The lines are: %d", count)
}
