// recover_divbyzero.go
package main

import (
	"fmt"
)

func badCall() {
  a, b := 10, 0
  n := a / b
  fmt.Println(n)
}

func test() {
   defer func() { 
    if e := recover(); e != nil {
       fmt.Printf("Panicing %s\r\n", e);
    }
   
    }()
    badCall()
    fmt.Printf("After bad call\r\n");
}

func main() {
   fmt.Printf("Calling test\r\n");
   test()
   fmt.Printf("Test completed\r\n");
}
/* Output:
Calling test
Panicing runtime error: integer divide by zero
Test completed
*/
