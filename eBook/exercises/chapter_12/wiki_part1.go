// wiki_part1.go
package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (this *Page) save() (err error) {
	return ioutil.WriteFile(this.Title, this.Body, 0666)
}

func (this *Page) load(title string) (err error) {
	this.Title = title
	this.Body, err = ioutil.ReadFile(this.Title)
	return err
}

func main() {
	page := Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()

	// load from Page.md
	var new_page Page
	new_page.load("Page.md")
	fmt.Println(string(new_page.Body))
}

/* Output:
 * # Page
 * ## Section1
 * This is section1.
 */
