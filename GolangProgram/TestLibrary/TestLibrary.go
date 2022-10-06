package testlibrary

import "fmt"

//define Employee struct
type Employee struct {
	FirstName string
	lastName  string
	Age       int
}

//define function setup lastname
func (e *Employee) SetLastName(lastName string) {
	e.lastName = lastName
}

//define function print lastname
func (e *Employee) PrintLastName() {
	fmt.Println("Last name is", e.lastName)
}
