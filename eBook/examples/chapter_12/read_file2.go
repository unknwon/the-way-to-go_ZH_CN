// read_csvfile.go
package main

import (
	"fmt"
	"os"
	//	"io/ioutil"
	//	"strings"
)

func main() {
	file, err := os.Open("products2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

/* Output:
[ABC FUNC GO]
[40 56 45]
[150 280 356]
*/
