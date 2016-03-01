package main

import "fmt"

type Day int

const (
    MO Day = iota
    TU
    WE
    TH
    FR
    SA
    SU
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func main() {
		var th Day = 3
		fmt.Printf("The 3rd day is: %s\n", th)
		// If index > 6: panic: runtime error: index out of range
		// but use the enumerated type to work with valid values:
		var day = SU;
		fmt.Println(day);  // prints Sunday
		fmt.Println(0, MO, 1, TU)
}
/* Output:
The 3rd day is: Thursday
Sunday
0 Monday 1 Tuesday
*/
