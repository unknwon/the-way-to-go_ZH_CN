// read input from the console:
package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin) // reader for input
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')

	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
