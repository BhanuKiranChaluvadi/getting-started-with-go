package main

import (
	"fmt"
	// 	"error"
)

func main() {
	port := 3000
	// p, err := startWebServer(port)
	// fmt.Println(p, err)
	_, err := startWebServer(port)
	fmt.Println(err)
}

// This declaration also works - port is assigned int b/c of next param
// two consecutive parameters are of same type.
// func startWebServer(port int, numberOfRetries int)
func startWebServer(port int) (int, error) {
	fmt.Println("Starting server ...")
	// do imported things
	fmt.Println("Server started on port", port)
	// return nil
	// return errors.New("Something went wrong")
	return port, nil
}
