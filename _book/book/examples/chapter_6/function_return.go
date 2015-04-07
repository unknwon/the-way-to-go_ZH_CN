package main
import "fmt"

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 2 gives: %v\n", p2(2))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(2))
}

func Add2() (func(b int) int) {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) (func(b int) int) {
	return func(b int) int {
		return a + b
	}
}
/* Output:
Call Add2 for 2 gives: 4
The result is: 4
*/