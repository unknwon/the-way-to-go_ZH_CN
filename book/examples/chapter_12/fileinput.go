package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// var inputFile *os.File
	// var inputError, readerError os.Error
	// var inputReader *bufio.Reader
	// var inputString string

	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" + 
		           "Does the file exist?\n" + 
			       "Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return // error or EOF
		}
		fmt.Printf("The input was: %s", inputString)
	}
}
