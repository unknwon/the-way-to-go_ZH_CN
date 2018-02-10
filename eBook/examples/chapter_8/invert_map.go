// invert_map.go
package main

import (
	"fmt"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
}

/* Output:
inverted:
Key: 12, Value: foxtrot / Key: 16, Value: hotel / Key: 87, Value: delta / Key: 23, Value: charlie /
Key: 65, Value: juliet / Key: 43, Value: kilo / Key: 56, Value: bravo / Key: 98, Value: lima /
Key: 34, Value: golf /
*/
