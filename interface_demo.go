package main

import (
	"fmt"
	"math"
)

/*
Interfaces are named collections of method signatures.
*/

type geometry interface {
	area() float64
	perim() float64
}

/*
struct are collections of properties
*/
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

/*
bind interface implementation to struct type
*/
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g, g.area(), g.perim())
}

func main() {
	r1 := &rect{
		width:  10.0,
		height: 10.0,
	}
	measure(r1)

	c1 := &circle{
		radius: 10.0,
	}

	measure(c1)
}
