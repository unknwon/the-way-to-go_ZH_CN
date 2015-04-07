package main

import "fmt"

type Base struct {
	id	string
}

func (b *Base) Id() string {
	return b.id
}

func (b *Base) SetId(id string) {
	b.id = id
}

type Person struct {
	Base
	FirstName		string
	LastName		string
}

type Employee struct {
	Person
	salary		float32
} 

func main() {
	idjb := Base{"007"}
	jb := Person{idjb, "James", "Bond"}
	e := &Employee{jb, 100000.}
	fmt.Printf("ID of our hero: %v\n", e.Id())
	// Change the id:
	e.SetId("007B")
	fmt.Printf("The new ID of our hero: %v\n", e.Id())
}
/* Output:
ID of our hero: 007
The new ID of our hero: 007B
*/