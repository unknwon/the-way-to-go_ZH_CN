// leaving out the length 4 in [4] uses  slice instead of an array
// which is generally better for performance
package main

import "fmt"

func main() {
	// var a = [4]float32 {1.0,2.0,3.0,4.0}
	var a = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("The sum of the array is: %f\n", Sum(a))
	var b = []int{1, 2, 3, 4, 5}
	sum, average := SumAndAverage(b)
	fmt.Printf("The sum of the array is: %d, and the average is: %f", sum, average)
}

/*
func Sum(a [4]float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}
*/

func Sum(a []float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}

func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum, float32(sum / len(a))
}
