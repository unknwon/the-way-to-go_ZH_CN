// json_xml_case.go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

type thing struct {
	Field1 int
	Field2 string
}

func main() {
	x := `<x><field1>423</field1><field2>hello from xml</field2></x>`
	j := `{"field1": 423, "field2": "hello from json"}`

	tx := thing{}
	if err := xml.Unmarshal(strings.NewReader(x), &tx); err != nil {
		log.Fatalf("Error unmarshaling XML: %v", err)
	}

	tj := thing{}
	if err := json.Unmarshal([]byte(j), &tj); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	fmt.Printf("From JSON: %#v\n", tj)
	fmt.Printf("From XML: %#v\n", tx)

}
/* Output with
type thing struct {
	Field1 int
	Field2 string
}:

From XML: main.thing{Field1:0, Field2:""}   // All matching is case sensitive! 
From JSON: main.thing{Field1:423, Field2:"hello from json"}

Output with
type thing struct {
	field1 int
	field2 string
}:

2012/02/22 10:51:11 Error unmarshaling JSON: json: cannot unmarshal object
field1" into unexported field field1 of type main.thing

JSON uses reflection to unmarshal!
*/