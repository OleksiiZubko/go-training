package main

import "fmt"

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

func initializeSlice() {
	primes := []int{2, 13, 17, 19, 23}
	fmt.Println(primes)
}

func appendSlice() {
	primes := []int{2, 3, 5, 7, 11}
	fmt.Println(len(primes), primes)

	primes = append(primes, 13)
	fmt.Println(len(primes), primes)

	primes = append(primes, 17, 19, 23)
	fmt.Println(len(primes), primes)
}

func subslice() {
	brothers := []string{"chico", "harpo", "groucho", "gummo", "zeppo"}
	fmt.Println(brothers)

	wellknown := brothers[0:3]
	fmt.Println(wellknown)

	a := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	fmt.Println(a)

	b := a[0:5]
	for i := 0; i < len(b); i++ {
		b[i] = b[i] * -1
	}
	fmt.Println(b)

	fmt.Println(a)
}

func bounds() {
	// primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	// fmt.Println(primes[-1])
	// fmt.Println("bounds: ", primes[10])
}

func main() {
	emptySlice()
	emptySlice20()
	zeroValues()
	nonEmptySlice()
	initializeSlice()
	appendSlice()
	subslice()
	bounds()
}
