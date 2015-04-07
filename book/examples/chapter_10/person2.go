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
}



