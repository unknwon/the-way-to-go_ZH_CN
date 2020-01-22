// pipeline1.go
package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("template test")
	t = template.Must(t.Parse("This is just static text. \n{{\"This is pipeline data - because it is evaluated within the double braces.\"}} {{`So is this, but within reverse quotes.`}}\n"))
	t.Execute(os.Stdout, nil)
}

/*
This is just static text.
This is pipeline data - because it is evaluated within the double braces. So is this, but within reverse quotes.
*/
