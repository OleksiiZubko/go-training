package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func countLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

func main() {
	lines, err := countLines(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
