package main

import "fmt"

func main() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
recover() // does not work here
	fmt.Println("Ending the program")
}