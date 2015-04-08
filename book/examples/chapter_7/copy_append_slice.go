package main

import "fmt"

func main() {
	sl_from := []int{1,2,3}
	sl_to := make([]int,10)
	
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)  // output: [1 2 3 0 0 0 0 0 0 0]
	fmt.Printf("Copied %d elements\n", n)  // n == 3
	
	sl3 := []int{1,2,3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)  // output: [1 2 3 4 5 6]
}
