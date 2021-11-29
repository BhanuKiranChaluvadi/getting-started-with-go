package main

import "fmt"

func main() {

	/// Example 1

	// when we declare constants
	// 1. we have to initialize them at the time we are declaring them (same line)
	// 2. The value of the constants has to be determined at compile time.
	//		We may tempted to assign the return value of a function to a constant
	// 		But thats not possible (functions are not evaluated until the runtime in go).
	const pi = 3.1415 // implicitly typed constant
	fmt.Println(pi)
	// pi = 4.13	// error

	/// Example 2

	// Implicitly typed constant. The compiler will interpret as appropriate
	// eveytime it runs into it.
	const c = 3

	fmt.Println(c + 3) // Data type of var C is dynamically interpreted as int

	// bunch of code

	// Here the data type of var C is dynamically interpreted as float32
	// since we did not mention any thing. Only floats add up but not integer + float
	fmt.Println(c + 1.2)

	/// Example 3
	const b int = 3
	fmt.Println(b + 3)
	// bunch of code
	// fmt.Println(b + 1.2) // error b/c int + float32
	fmt.Println(float32(b) + 1.2)

}
