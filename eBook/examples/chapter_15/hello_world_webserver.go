// hello_world_webserver.go
package main

import (
	"fmt"
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	//fmt.Fprint(w, "Hello, " + req.URL.Path[1:])
	fmt.Fprintf(w, "Hello, %s ", req.URL.Path[1:])
	// io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
	// http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))
}
