package main

import "fmt"

func main() {
	var firstName *string
	fmt.Println(firstName)

	// This assignment not going to wrok because we need to assign memory first.
	// And the pointer is unintialized.
	// *firstName = "BhanuKiran"

	var secondName *string = new(string)
	*secondName = "BhanuKiran"
	fmt.Println(secondName)
	fmt.Println(*secondName)

	thirdName := "Ragini"

	ptr := &thirdName
	fmt.Println(ptr, *ptr)

	thirdName = "Sravya"
	fmt.Println(ptr, *ptr)
}
