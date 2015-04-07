// Q28_word_letter_count.go
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
var nrchars, nrwords, nrlines int

func main() {
	nrchars, nrwords, nrlines = 0, 0, 0 
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop: ")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		if input == "S\r\n" { // Windows, on Linux it is "S\n"
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Counters(input)
	}
}

func Counters(input string) {
	nrchars += len(input) - 2 // -2 for \r\n
	// count number of spaces, nr of words is +1
	nrwords += strings.Count(input, " ") + 1
	nrlines++
}
