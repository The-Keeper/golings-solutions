// structs2
// Make me compile!

package main

import "fmt"


type ContactDetails struct {
	phone string
}

type Person struct {
	ContactDetails
	name string
	age int
}


func main() {
	contactDetails := ContactDetails{phone: "12-34-56"}
	person := Person{name: "John", age: 32}
	person.phone = contactDetails.phone
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
