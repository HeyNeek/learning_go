package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	nick := person{
		firstName: "Nick",
		lastName: "Aguirre",
		contact: contactInfo{
			email: "nickpwn@live.com",
			zipCode: 97201,
		},
	}

	fmt.Println(nick)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

