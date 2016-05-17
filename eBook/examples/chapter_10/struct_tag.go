package main

import (
	"fmt"
	"reflect"
)

type TagType1 struct { //  tags
	field1 bool   // "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType1{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType1, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
	fmt.Printf("%v\n", ixField)
}

/* Output:
An important answer
The name of the thing
How much there are
*/
