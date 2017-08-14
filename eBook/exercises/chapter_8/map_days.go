// map_days.go
package main

import (
	"fmt"
)

var Days = map[int]string{1: "monday",
	2: "tuesday",
	3: "wednesday",
	4: "thursday",
	5: "friday",
	6: "saturday",
	7: "sunday"}

func main() {
	fmt.Println(Days)
	// fmt.Printf("%v", Days)
	flagHolliday := false
	for k, v := range Days {
		if v == "thursday" || v == "holliday" {
			fmt.Println(v, " is the ", k, "th day in the week")
			if v == "holliday" {
				flagHolliday = true
			}
		}
	}
	if !flagHolliday {
		fmt.Println("holliday is not a day!")
	}
}

/* Output:
thursday  is the  4 th day in the week
holliday is not a day!
*/
