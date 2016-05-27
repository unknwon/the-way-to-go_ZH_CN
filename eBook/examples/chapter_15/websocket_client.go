// websocket_client.go
package main

import (
	"fmt"
	"time"
	"./src/golang.org/x/net/websocket"
)

func main() {
println("Dial")
	ws, err := websocket.Dial("ws://localhost:12345/websocket", "",
		"http://localhost/")
	if err != nil {
		panic("Dial: " + err.Error())
	}
	go readFromServer(ws)
	time.Sleep(5e9)
    ws.Close()
}

func readFromServer(ws *websocket.Conn) {
	buf := make([]byte, 1000)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}
	}
}
