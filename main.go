package main

import "fmt"

func multipleAssigments() {
	// you can declare multiple variables at the same time.
	var x, y, z = 1, 2, 3
	fmt.Println(x, y, z)

	// they don't even need to be the same type.
	var a, b, c = 1, "two", 3.3
	fmt.Println(a, b, c)

	// the short declaration syntax also works.
	i, j, k := 7, 11, 13
	fmt.Println(i, j, k)

	// you can use type conversions to ensure the type of
	// the value is the one you wish.
	d, e, f := uint16(20), int8(99), uint(12)

	// let's check that d, e, and f are the types we expect
	var g uint16 = d // would fail if d was not of type uint16
	var h int8 = e
	var l uint = f

	fmt.Println(g, h, l)
}

func loop() {
	var i = 1
	for {
		if i > 10 {
			break
		}
		if i%2 == 0 {
			println(i)
		}
		i++
	}
}

// A Double returns a number two times larger than i.
func A(i int) int {
	return i * 2
}

// multipleReturnValues switches the values of a and b
func multipleReturnValues(a int, b int) (int, int) {
	return b, a
}

func f() (int, int, int, bool, string) {
	return 0, 1, 2, true, "test"
}

func main() {
	loop()
	multipleAssigments()

	fmt.Println(A(200))

	a, b := multipleReturnValues(7, 9)
	fmt.Println(a, b)

	a, _, _, c, _ := f()
	fmt.Println(a, c)
}
