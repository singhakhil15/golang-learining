package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alexA := person{firstName: "AlexAAA", lastName: "Anderson"}
	// var alexB person
	// alexB.firstName = "AlexBBBB"
	// alexB.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Println(alexA)
	// fmt.Println(alexB)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@jim.com",
			zipCode: 3456,
		},
	}

	fmt.Println(jim)
	fmt.Printf("%+v", jim)

	jim.updateName("jimmy")
	jim.print()

	// & is oprater used to get access of varibles memory address.
	jimPointer := &jim
	jimPointer.updateNamePointer("jimmy")
	jimPointer.print()

	jim.updateNamePointer(" Akhil ")
	jim.print()
}

func (p person) print() {
	fmt.Printf("\n %+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// * oprater used to get value of memory address is pointing at
func (p *person) updateNamePointer(newFirstName string) {
	(*p).firstName = newFirstName
}
