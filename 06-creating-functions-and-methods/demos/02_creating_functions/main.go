package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	port := 3000
	startWebServer(port, 2)
}

// This declaration also works - port is assigned int b/c of next param
// two consecutive parameters are of same type.
// func startWebServer(port int, numberOfRetries int)
func startWebServer(port, numberOfRetries int) {
	fmt.Println("Starting server ...")
	// do imported things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}
