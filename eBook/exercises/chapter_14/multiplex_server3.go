// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import "fmt"

type Request struct {
	a, b   int
	replyc chan int // reply channel inside the Request
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func (r *Request) String() string {
	return fmt.Sprintf("%d+%d=%d", r.a, r.b, <-r.replyc)
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return // stop infinite loop
		}
	}
}

func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func main() {
	adder, quit := startServer(func(a, b int) int { return a + b })
	// make requests:
	req1 := &Request{3, 4, make(chan int)}
	req2 := &Request{150, 250, make(chan int)}
	// send requests on the service channel
	adder <- req1
	adder <- req2
	// ask for the results: ( method String() is called )
	fmt.Println(req1, req2)
	// shutdown server:
	quit <- true
	fmt.Print("done")
}

/* output:
3+4=7 150+250=400
done
*/
