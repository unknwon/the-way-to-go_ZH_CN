// poll_url.go
package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

func main() {
	// Execute an HTTP HEAD request for all url's 
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error", url, err)
		}
		fmt.Print(url, ": ", resp.Status)
	}
}
/* Output:
http://www.google.com/ :  302 Found
http://golang.org/ :  200 OK
http://blog.golang.org/ :  200 OK
*/
