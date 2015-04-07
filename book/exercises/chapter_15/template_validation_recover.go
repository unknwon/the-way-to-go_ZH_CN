// template_validation_recover.go
package main

import (
	"text/template"
	"fmt"
	"log"
)

func main() {
	tOk := template.New("ok")
	tErr := template.New("error_template")
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()

	//a valid template, so no panic with Must:
	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}")) 
	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}
/* Output:
The first one parsed OK.
The next one ought to fail.
2011/10/27 10:56:27 run time panic: template: error_template:1: unexpected "}" in command
*/