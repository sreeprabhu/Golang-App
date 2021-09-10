package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	per1 := person{"Sreeprabhu", "PA", contactInfo{"", ""}}
	per2 := person{firstName: "Sreeprabhu", lastName: "PA"}
	fmt.Println("Sreeprabhu: ", per1, per2)

	// initialize a struct
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	//nested struct
	jim := person{
		firstName: "Jim",
		lastName:  "John",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "679337",
		},
	}
	// fmt.Printf("%+v", jim)

	// jim.updateName("jimmy")
	// update this code with the below code

	// jimPointer := &jim // `&bvariable` gives the memory address of the value, this variable is pointing to
	// jimPointer.updateName("jimmy")

	// applying back the old code, as go will automatically map the address of jim and pass to the updateName fun
	jim.updateName("jimmy")

	jim.print()
}

/**
@func updateName - this function can only be called with a receiver of a pointer to a person struct
`*pointer` - this is a type description(not a pointer as such), it means we are working with a pointer to a person
*/
func (pointerToPerson *person) updateName(newFirstName string) {
	// Go is pass by value language
	// here when the fun updateName receives the struct person, Go will create a copy of the person,
	// in a different location in our system RAM, so instead of updating the original value,
	// this func would update the copy of the struct person
	// so we are using pointers to overcome this issue

	// *pointerToPerson - This is an operator, it means we want to manipulate the value, the pointer is referencing
	(*pointerToPerson).firstName = newFirstName // `*address` gives the value this memory address is pointing at
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
