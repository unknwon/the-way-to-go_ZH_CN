// Concurrent computation of pi.
// See http://goo.gl/ZuTZM. - Comparison with Scala!
//
// This demonstrates Go's ability to handle
// large numbers of concurrent processes.
// It is an unreasonable way to calculate pi.
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(CalculatePi(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta) 
}

// CalculatePi launches n goroutines to compute an
// series-approximation of pi.
func CalculatePi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		// calculate k-th term in the series
		go term(ch, float64(k))
	}
	f := 0.0
	//wait for all goroutines to complete, get and sum up their results:
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
/* Output:
3.14179261359579
The calculation took this amount of time: 0.028002
*/