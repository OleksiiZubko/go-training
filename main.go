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

func functions() {
	fmt.Println(A(200))

	a, b := multipleReturnValues(7, 9)
	fmt.Println(a, b)

	a, _, _, c, _ := f()
	fmt.Println(a, c)
}

func maps() {
	m := make(map[string]int)
	fmt.Println(m, len(m))

	days := make(map[int]string)
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	days[4] = "Thursday"
	days[5] = "Friday"
	days[6] = "Saturday"
	days[7] = "Sunday"
	fmt.Println(days)
	fmt.Println("The 3th day is: ", days[3])

	days2 := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	fmt.Println(days2)

	// map of planets to their number of moons
	moons := map[string]int{
		"Mercury": 0,
		"Venus":   0,
		"Earth":   1,
		"Mars":    2,
		"Jupiter": 67,
	}

	fmt.Println("Earth:", moons["Earth"])
	fmt.Println("Neptune:", moons["Neptune"])

	moons2 := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

	n, found := moons2["Earth"]
	fmt.Println("Earth:", n, "Found:", found)

	n, found = moons2["Neptune"]
	fmt.Println("Neptune:", n, "Found:", found)

	n, found = moons2["Mercury"]
	fmt.Println("Mercury:", n, "Found:", found)

	moons3 := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

	const planet = "Mars"

	n, found = moons3[planet]
	fmt.Println(planet, n, "Found:", found)

	delete(moons3, planet)

	n, found = moons3[planet]
	fmt.Println(planet, n, "Found:", found)

	cities := map[string]int{
		"Tokyo": 33200000, "New York": 17800000,
		"Sau Paulo": 17700000, "Delhi": 14300000,
		"Moscow": 10500000,
	}

	for name, pop := range cities {
		fmt.Println("City:", name, "Population", pop)
	}
}

func slices() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for i, p := range primes {
		fmt.Println("The", i, "th prime is", p)
	}
}

func switchF() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for i, p := range primes {
		i++
		switch i {
		case 1:
			fmt.Println("The", i, "st prime is", p)
		case 2:
			fmt.Println("The", i, "nd prime is", p)
		case 3:
			fmt.Println("The", i, "rd prime is", p)
		default:
			fmt.Println("The", i, "th prime is", p)
		}
	}
}

func fmtF() {
	name := "David"
	age := 27

	fmt.Printf("Hello my name is %s, my age is %d", name, age)
}

func fmtF2() {
	name := "David"
	age := 27

	fmt.Printf("Hello my name is %d, my age is %d\n", name)

	_ = age // keep the compiler happy
}

func fmtF3() {
	name := "David"
	age := 27

	fmt.Printf("Hello my name is %v, my age is %v", name, age)
}

func scopeF1(x int) {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}

func scopeF() {
	var x = 200
	scopeF1(x)
}

func main() {
	loop()
	multipleAssigments()
	functions()
	maps()
	slices()
	switchF()
	fmtF()
	fmtF2()
	fmtF3()
	scopeF()
}
