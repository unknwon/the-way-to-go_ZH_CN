// template_field.go
package main

import (
	"fmt"
    "os"
    "text/template"
)

type Person struct {
    Name string
    nonExportedAgeField string 
}

func main() {
    t := template.New("hello") 
    t, _ = t.Parse("hello {{.Name}}!") 
//    t, _ = t.Parse("hello {{.Name}}!  {{.nonExportedAgeField}} ")  // There was an error: template: hello:1:20: executing "hello" at <.nonExportedAgeField>: nonExportedAgeField is an unexported field of struct type main.Person
    p := Person{Name:"Mary", nonExportedAgeField: "31"}
    if err := t.Execute(os.Stdout, p); err != nil {
        fmt.Println("There was an error:", err.Error())
    }
}
// Output:   hello Mary!