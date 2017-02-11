package main

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b   int
	replyc chan int
}

func process(r *Request) {
	// do something
}

func handle(r *Request) {
	sem <- 1 // doesn't matter what we put in it
	process(r)
	<-sem // one empty place in the buffer: the next request can start
}

func server(service chan *Request) {
	for {
		request := <-service
		go handle(request)
	}
}

func main() {
	service := make(chan *Request)
	go server(service)
}
