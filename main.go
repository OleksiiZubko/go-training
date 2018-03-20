package main

func main() {
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
