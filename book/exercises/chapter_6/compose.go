// compose.go
package main

import (
	"fmt"
	"math"
)

func Compose(f, g func(x float64) float64) func(x float64) float64 {
	return func(x float64) float64 { // closure
		return f(g(x))
	}
}

func main() {
	fmt.Print(Compose(math.Sin, math.Cos)(0.5)) // output: 0.7691963548410085
}
