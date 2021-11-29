package main

import (
	"fmt"

	"github.com/BhanuKiranChaluvadi/getting-started-with-go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Bhanu Kiran",
		LastName:  "Chaluvadi",
	}
	fmt.Println(u)
}
