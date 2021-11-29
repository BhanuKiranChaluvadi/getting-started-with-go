package main

import "fmt"

func main() {

	/// Example 1
	arr := [3]int{1, 2, 3}
	slice := arr[:] // slice is pointing to the data that the array is keeping

	arr[1] = 42   // identical data
	slice[2] = 27 // identical data

	fmt.Println(arr)
	fmt.Println(slice) // identical data

	/// Example 2
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	slice2 = append(slice2, 4)
	fmt.Println(slice2)

	slice2 = append(slice2, 42, 27)
	fmt.Println(slice2)

	slice3 := slice2[1:]
	slice4 := slice2[:2]
	slice5 := slice2[1:2]

	fmt.Println(slice3, slice4, slice5)

	slice6 := []int{5, 6, 7}
	slice7 := append(slice2, slice6...)
	fmt.Println(slice7)

}
