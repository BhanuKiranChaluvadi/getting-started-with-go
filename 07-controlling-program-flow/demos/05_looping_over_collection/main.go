package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	// uglier
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		fmt.Println(k, v)
	}

	for k := range wellKnownPorts {
		fmt.Println(k)
	}

	for _, v := range wellKnownPorts {
		fmt.Println(v)
	}

}
