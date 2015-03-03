// panic_package.go
package main

import (
	"fmt"
	"./parse/parse"
)

func main() {
        var examples = []string{
                "1 2 3 4 5",
                "100 50 25 12.5 6.25",
                "2 + 2 = 4",
                "1st class",
                "",
        }

        for _, ex := range examples {
                fmt.Printf("Parsing %q:\n  ", ex)
                nums, err := parse.Parse(ex)
                if err != nil {
                        fmt.Println(err) // here String() method from ParseError is used
                        continue
                }
                fmt.Println(nums)
        }
}
/* Output:
Parsing "1 2 3 4 5":
  [1 2 3 4 5]
Parsing "100 50 25 12.5 6.25":
  pkg parse: error parsing "12.5" as int
Parsing "2 + 2 = 4":
  pkg parse: error parsing "+" as int
Parsing "1st class":
  pkg parse: error parsing "1st" as int
Parsing "":
  pkg: no words to parse
*/
