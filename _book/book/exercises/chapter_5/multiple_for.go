// multiple_for.go
package main

import "fmt"

func main() {
    //multiple initialization; a consolidated bool expression with && and ||; multiple ‘incrementation’
    for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, 
		j+1, s + "a"  {
        fmt.Println("Value of i, j, s:", i, j, s)
    }
}
/* Output:
Value of i, j, s: 0 5 a
Value of i, j, s: 1 6 aa
Value of i, j, s: 2 7 aaa
*/