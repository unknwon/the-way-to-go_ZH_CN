// panic_defer_convint.go
package main

import (
	"fmt"
	"math"
)

func main() {
	l := int64(15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}
	fmt.Println()
	l = int64(math.MaxInt32 + 15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d", l, i)
	}
}

func ConvertInt64ToInt(l int64) int {
	if math.MinInt32 <= l && l <= math.MaxInt32 {
		return int(l)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", l))
}

func IntFromInt64(l int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = ConvertInt64ToInt(l)
	return i, nil
}

/* Output:
15000 converted to an int32 is 15000
The conversion of 2147498647 to an int32 resulted in an error: 2147498647 is out of the int32 range
*/
