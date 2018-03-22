package main

import "fmt"

func main() {
	emptySlice()
	emptySlice20()
	zeroValues()
	nonEmptySlice()
}

func emptySlice() {
	var i []int
	fmt.Println(len(i))
}

func emptySlice20() {
	i := make([]int, 20)
	fmt.Println(len(i))
}

func zeroValues() {
	x := make([]int, 5)
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

func nonEmptySlice() {
	primes := make([]int, 10)
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	primes[3] = 7
	primes[4] = 11
	primes[5] = 13
	primes[6] = 17
	primes[7] = 19
	primes[8] = 23
	primes[9] = 29
	fmt.Println(primes)
}
