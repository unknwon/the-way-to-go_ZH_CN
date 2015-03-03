package main

import "fmt"

func main() {
	str := "Google"
	str2 := Split2(str)
	fmt.Printf("The string %s transformed is: %s\n", str, str2)
}

func Split2(s string) string {
	mid := len(s) / 2
	return s[mid:] + s[:mid]
}
// Output: The string Google transformed is: gleGoo