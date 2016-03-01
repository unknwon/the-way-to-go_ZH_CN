package main
import "fmt"

func main() {
	var n int16 = 34
	var m int32
	
// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n 
	m = int32(n)
	
	fmt.Printf("32 bit int is:  %d\n", m)
	fmt.Printf("16 bit int is:  %d\n", n)
}
/* Output:
32 bit int is:  34
16 bit int is:  34
*/