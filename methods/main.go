package main

import (
	"fmt"
)

// Point - the coordinates of the point
type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("point: x=%d, y=%d", p.X, p.Y)
}

func main() {
	p := Point{X: 300, Y: 60}

	println(p.String())
}
