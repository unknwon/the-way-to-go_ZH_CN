package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	for i:=0; i <= 25; i++ {
		result = fibonacci(i) 
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta) 
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
/* Output:
fibonacci(0) is: 1
fibonacci(1) is: 1
fibonacci(2) is: 2
fibonacci(3) is: 3
fibonacci(4) is: 5
fibonacci(5) is: 8
fibonacci(6) is: 13
fibonacci(7) is: 21
fibonacci(8) is: 34
fibonacci(9) is: 55
fibonacci(10) is: 89
fibonacci(11) is: 144
fibonacci(12) is: 233
fibonacci(13) is: 377
fibonacci(14) is: 610
fibonacci(15) is: 987
fibonacci(16) is: 1597
fibonacci(17) is: 2584
fibonacci(18) is: 4181
fibonacci(19) is: 6765
fibonacci(20) is: 10946
fibonacci(21) is: 17711
fibonacci(22) is: 28657
fibonacci(23) is: 46368
fibonacci(24) is: 75025
fibonacci(25) is: 121393
longCalculation took this amount of time: 3.0001ms
*/
