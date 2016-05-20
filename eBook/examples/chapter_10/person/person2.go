package person

type Person struct {
	firstName	string
	lastName	string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	 p.firstName = newName

// test
	p.lastName = p.firstName + "X"
}

// make lastName readonly by outside
func (this *Person) LastName() string{
	return this.lastName
}

