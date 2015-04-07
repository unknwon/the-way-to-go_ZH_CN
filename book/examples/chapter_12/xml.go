// xml.go
package main

import (
	"fmt"
	"strings"
	"encoding/xml"
)

var t, token	xml.Token
var	err			error

func main() {
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	
	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
			case xml.StartElement:
				name := token.Name.Local
				fmt.Printf("Token name: %s\n", name)
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
					// ...
				}
			case xml.EndElement:
				fmt.Println("End of token")
			case xml.CharData:
				content := string([]byte(token))
				fmt.Printf("This is the content: %v\n", content )
				// ...
			default:
				// ...
		}
	}
}
/* Output:
Token name: Person
Token name: FirstName
This is the content: Laura
End of token
Token name: LastName
This is the content: Lynn
End of token
End of token
*/
