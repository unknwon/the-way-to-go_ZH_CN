package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
fmt.Println("starting")
	err := http.ListenAndServe("localhost:8080", nil)
//fmt.Println("started") // not printed
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
