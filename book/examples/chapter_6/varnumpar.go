package main

import "fmt"

func main() {
	x := Min(1, 3, 2, 0) 
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7,9,3,5,1}
	x = Min(arr...) 
	fmt.Printf("The minimum in the array arr is: %d", x)
}

func Min(a ...int) int {
  if len(a)==0 {
    return 0
  }
  min := a[0]
  for _, v := range a {
    if v < min {
      min = v
    }
  }
  return min
}
/*
The minimum is: 0
The minimum in the array arr is: 1
*/