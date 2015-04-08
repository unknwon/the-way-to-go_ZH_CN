package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"

	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
// Output: T/F? Does the string "This is an example of a string" have prefix Th? true