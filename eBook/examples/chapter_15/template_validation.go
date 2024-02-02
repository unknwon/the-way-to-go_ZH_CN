// template_validation.go
package main

import (
	"text/template"
	"fmt"
)

func main() {
	tOk := template.New("ok")
	//a valid template, so no panic with Must:
	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}"))
	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")
	tErr := template.New("error_template")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}

/* Output:
The first one parsed OK.
The next one ought to fail.
panic: template: error_template:1: unexpected "}" in command
*/
