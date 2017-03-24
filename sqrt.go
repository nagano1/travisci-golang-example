package main

import (
	"fmt"
)

// Inte is greatest interface
type Inte interface {
	afbbw() int
}

// Again is not stricts
type Again struct {
	a, x int
	y, g string
}

// Test tests
func (a Again) Test(x int) int {
	return a.a
}

func (a Again) afbbw() int {
	return 5
}

// Sqrt goes wrong
func Sqrt(x float64) float64 {
	z := 0.3
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

func main() {
	var a Inte

	fmt.Println("fwefw")
	aga := Again{y: "foiew", a: 32, x: 54, g: "str"}
	b := aga.Test(3)
	a = aga
	if a != nil {

	}
	fmt.Print(b)
}
