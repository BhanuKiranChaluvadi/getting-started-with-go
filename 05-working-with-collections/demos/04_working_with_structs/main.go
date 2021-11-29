package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Bhanu Kiran"
	u.LastName = "Chaluvadi"
	fmt.Println(u)
	fmt.Println(u.ID)

	u2 := user{ID: 1,
		FirstName: "Bhanu Kiran",
		LastName:  "Chaluvadi"}
	fmt.Println(u2)
}
