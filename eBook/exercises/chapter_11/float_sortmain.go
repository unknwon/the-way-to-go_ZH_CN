package main

import (
	"./float64"
	"fmt"
)

func main() {
	f1 := float64.NewFloat64Array()
	f1.Fill(10)
	fmt.Printf("Before sorting %s\n", f1)
	float64.Sort(f1)
	fmt.Printf("After sorting %s\n", f1)
	if float64.IsSorted(f1) {
		fmt.Println("The float64 array is sorted!")
	} else {
		fmt.Println("The float64 array is NOT sorted!")
	}
}

/* Output:
Before sorting { 55.0 82.3 36.4 66.6 25.3 82.7 47.4 21.5 4.6 81.6  }
After sorting { 4.6 21.5 25.3 36.4 47.4 55.0 66.6 81.6 82.3 82.7  }
The float64 array is sorted!
*/
