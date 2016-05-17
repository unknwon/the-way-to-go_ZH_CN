// method_on_time.go
package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time //anonymous field
	a         int
}

func (t *myTime) first3Chars() string {
	t.a += 1
	return t.Time.String()[0:3]
}

func (t *myTime) first() string {
	return "a"
}

func main() {
	m := myTime{time.Now(), 3}
	fmt.Printf("%v\n", m.a)
	println(m.first())
	fmt.Println("Full time now:", m.String())      //calling existing String method on anonymous Time field
	fmt.Println("First 3 chars:", m.first3Chars()) //calling myTime.first3Chars
	fmt.Println(m.a)
}

/* Output:
Full time now: Mon Oct 24 15:34:54 Romance Daylight Time 2011
First 3 chars: Mon
*/
