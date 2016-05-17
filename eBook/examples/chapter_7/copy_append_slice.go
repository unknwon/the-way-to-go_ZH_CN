package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)                    // output: [1 2 3 0 0 0 0 0 0 0]
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	// output: [1 2 3 4 5 6]

	sl3 = append(sl3, sl_from...)
	fmt.Println(sl3)

	fmt.Printf("len %d cap %d", len(sl3), cap(sl3))

	//	sl3 = AppendByte(sl3, sl_from)
	sl3 = AppendByte(sl3, sl3)
	fmt.Println(sl3)
	fmt.Printf("len %d cap %d", len(sl3), cap(sl3))
}

func AppendByte(slice []int, data []int) []int {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]int, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
