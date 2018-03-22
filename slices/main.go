package main

import "fmt"

func main() {
	emptySlice()
	emptySlice20()
}

func emptySlice() {
	var i []int
	fmt.Println(len(i))
}

func emptySlice20() {
	i := make([]int, 20)
	fmt.Println(len(i))
}
