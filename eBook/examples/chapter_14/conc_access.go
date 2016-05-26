// conc_access.go
package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend() {
fmt.Println("backend begin")
	for f := range p.chF {
fmt.Printf(" call f %#v\n ", f)
		f()
	}
fmt.Println("backend end")
}

// Set salary.
func (p *Person) SetSalary(sal float64) {
fmt.Println("set begin")
	p.chF <- func() { p.salary = sal }
fmt.Println("set end")
}

// Retrieve salary.
func (p *Person) Salary() float64 {
fmt.Println("get begin")
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }
fmt.Println("get end")
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	bs := NewPerson("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed:")
	fmt.Println(bs)
}
/* Output:
Person - name is: Smith Bill - salary is: 2500.50
Salary changed:
Person - name is: Smith Bill - salary is: 4000.25
*/
