package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	// SIMPLE STRUCT
	// alex := person{firstName: "Alex", lastName: "Anderson", age: 35}

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.age = 35

	// fmt.Printf("%+v\n", alex)

	// NESTED STRUCTS
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		age:       44,
		contact: contactInfo{
			email:   "jimboy@hotmail.com",
			zipcode: 93800,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Johnny")
	jim.updateName("Johnny")
	jim.print()
}

// Use pointers to pass values by reference and update `jim` directly
func (ptrToPerson *person) updateName(newFirstName string) {
	(*ptrToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
