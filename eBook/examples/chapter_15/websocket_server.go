// websocket_server.go
package main

import (
	"fmt"
	"net/http"
	"./src/golang.org/x/net/websocket"
)

func server(ws *websocket.Conn) {
	fmt.Printf("new connection\n")
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
	}
	fmt.Printf(" => closing connection\n")
	ws.Close()
}

func main() {
println("Handle")
	http.Handle("/websocket", websocket.Handler(server))
println("ListenAndServe")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
