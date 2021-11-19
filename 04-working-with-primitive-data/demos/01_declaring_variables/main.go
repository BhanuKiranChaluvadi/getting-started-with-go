package main

import "fmt"

func main() {
	// declare
	var i int
	// initialize
	i = 143
	fmt.Println(i)

	//  declare & initialize - float 32/64 bits
	var f float32 = 0.143
	fmt.Println(f)

	// Not starting with "var" keyword
	// : colon before the equal sign. This is going to allow us to use
	// implicit initilization syntax (imply the datatype based on assignment)
	firstName := "BhanuKiran"
	fmt.Println(firstName)

	// boolean
	b := true
	fmt.Println(b)

	// complex number
	c := complex(3, 4)
	fmt.Println(c)

	// multiple assignment in single line
	// split the complex number
	r, im := real(c), imag(c)
	fmt.Println(r, im)

}
