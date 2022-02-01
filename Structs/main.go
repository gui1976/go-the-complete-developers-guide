package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex.anderson@whatever.com"
	alex.contact.zipCode = 1234

	fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@whatever.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", jim)
}
