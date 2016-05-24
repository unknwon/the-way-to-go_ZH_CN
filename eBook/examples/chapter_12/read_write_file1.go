// read_write_file.go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "products.txt"
//	inputFile := "/media/sf_D_DRIVE/software/YLMF_GHOST_WIN8_X64.iso" // will we OOM? yes: fatal error: runtime: out of memory
	outputFile := "products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
fmt.Println(len(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
