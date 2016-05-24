// degob.go
package main

import (
	"bufio"
	"fmt"
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type             string
	City             string
	Country          string
}

type VCard struct {
	FirstName	string
	LastName	string
	Addresses	[]*Address
	Remark		string
}

var content	string
var vc VCard

func main() {
		// using a decoder:
	file, _ := os.Open("vcard.gob")
	defer file.Close()
	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc, vc.Addresses[0], vc.Addresses[1])
}
// Output:
// {Jan Kersschot [0x12642e60 0x12642e80] none}