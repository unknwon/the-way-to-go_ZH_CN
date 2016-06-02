package main

import "fmt"
import "time"
import "net/http"
import _ "net/http/pprof"

/*
Test golang performance monitoring.
See http://blog.ralch.com/tutorial/golang-performance-and-memory-analysis/
*/
func main() {

	go func() {
		fmt.Println("starting")
		http.ListenAndServe(":8080", http.DefaultServeMux)
	}()

	for i := 0; i < 999; i++ {
		fmt.Printf("The counter is at %d\n", i)
		time.Sleep(2 * time.Second)
		//		time.Sleep(1e9)
	}
}
