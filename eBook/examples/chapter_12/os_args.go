// os_args.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice "
fmt.Printf("%v\n", os.Args)
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
