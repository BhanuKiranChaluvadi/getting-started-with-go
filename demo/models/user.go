package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1 // Notice that we don´t have a datatype in here
	// Because we have this initilized at the package level
	// Don´t need colon : operator
)
