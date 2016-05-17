// big.go
package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Here are some calculations with bigInts:
	im := big.NewInt(2)
	in := im
	io := big.NewInt(3)
	ip := big.NewInt(4)
	fmt.Printf("ip %v\n", ip)
	//	ip.Mul(im, io)
	ip.Add(im, io)
	//	ip.Div(ip, im)
	//	ip.Sub(ip, im)
	fmt.Printf("ip %v\n", ip)

	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	// Here are some calculations with bigInts:
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
}

/* Output:
Big Int: 43492122561469640008497075573153004
Big Rat: -37/112
*/
