// helloworld.go
package helloworld

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>Hello, World! 세상아 안녕!! </body></html>")
}

