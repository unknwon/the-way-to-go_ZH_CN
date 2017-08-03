package main

import "fmt"

func main() {
	printrec(1)
}

func printrec(i int) {
	if i > 10 {
		return
	}
	printrec(i + 1)
	fmt.Printf("%d ", i)
}
