package main

func main() {
	var i int

	// ugly infinite loop
	for {
		// if we don´t have this condition, its an infinite
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}
