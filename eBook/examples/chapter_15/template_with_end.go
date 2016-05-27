// template_with_end.go
package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")
	t.Execute(os.Stdout, nil)

	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n") 
	t.Execute(os.Stdout, nil)
}
/* Output:
hello!
hello Mary!
*/