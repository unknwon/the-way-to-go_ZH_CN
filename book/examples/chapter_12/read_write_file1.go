// read_write_file.go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "products.txt"
	outputFile := "products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x644)
	if err != nil {
		panic(err.Error())
	}
}
