package main

// import block
import (
	"fmt"
)

// constants at the package level
// This constant is now accessible throughout the entire pkg
// Its not limited or scoped to the main function.
const pi = 3.1415

// iota starts with zero and increments every single resue
const (
	sixth   = iota
	seventh = iota
)

// iota reset to "zero" in each const block
// It automatically assing the previous expression to next
// variable if not defined.
const (
	eighth = iota
	nineth
)

// const block - much like import block
const (
	f = iota + 6  // 0 + 6
	s = 2 << iota //  2 << 1  ( 2 * 2  = 4)
)

//evaluate previous expression
const (
	a = iota + 7 // 0 + 7
	b            // evalute previous expression with incremented iota b = 1 + 7
)

const (
	first  = 1
	second = "second"
	third  = "third"
	fourth = iota + 11 // special keyword iota
	fifth  = iota
)

func main() {
	fmt.Println(pi)
	fmt.Println(first, second)
	fmt.Println(third)
	fmt.Println(fourth, fifth)
	fmt.Println(sixth, seventh, eighth, nineth)
	fmt.Println(f, s)
	fmt.Println(a, b)
}
