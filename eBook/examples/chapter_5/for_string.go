// for_string.go
package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
}

/* Output:
The length of str is: 27
Character on position 0 is: G
Character on position 1 is: o
Character on position 2 is:
Character on position 3 is: i
Character on position 4 is: s
Character on position 5 is:
Character on position 6 is: a
Character on position 7 is:
Character on position 8 is: b
Character on position 9 is: e
Character on position 10 is: a
Character on position 11 is: u
Character on position 12 is: t
Character on position 13 is: i
Character on position 14 is: f
Character on position 15 is: u
Character on position 16 is: l
Character on position 17 is:
Character on position 18 is: l
Character on position 19 is: a
Character on position 20 is: n
Character on position 21 is: g
Character on position 22 is: u
Character on position 23 is: a
Character on position 24 is: g
Character on position 25 is: e
Character on position 26 is: !
The length of str2 is: 9
Character on position 0 is: æ
Character on position 1 is: 
Character on position 2 is: ¥
Character on position 3 is: æ
Character on position 4 is: 
Character on position 5 is: ¬
Character on position 6 is: è
Character on position 7 is: ª
Character on position 8 is: 
*/
