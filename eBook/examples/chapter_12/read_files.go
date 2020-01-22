// read_files.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Printf("Reading files...\n")
	flag.Parse()

	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("[File: %v]\n", flag.Arg(i))
		fin, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Printf("The file %v does not exist!\n", flag.Arg(i))
			break
		}
		r := bufio.NewReader(fin)
		for line, _, err := r.ReadLine(); err != io.EOF; line, _, err = r.ReadLine() {
			fmt.Printf("Lines: %v (error %v)\n", string(line), err)
		}
		fin.Close()
	}
}
