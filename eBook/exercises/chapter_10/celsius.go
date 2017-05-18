// celsius.go
package main

import (
	"fmt"
	"strconv"
)

type Celsius float64

func (c Celsius) String() string {
	return "The temperature is: " + strconv.FormatFloat(float64(c), 'f', 1, 32) + " °C"
}

func main() {
	var c Celsius = 18.36
	fmt.Println(c)
}

// The temperature is: 18.4 °C
