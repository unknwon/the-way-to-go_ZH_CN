// rpc_server.go
// after client-exits the server shows the message:
//       1:1234: The specified network name is no longer available.
//       2011/08/01 16:19:04 rpc: rpc: server cannot decode request: WSARecv tcp 127.0.0.
package main

import (
	"net/http"
	"log"
	"net"
	"net/rpc"
	"time"
	"./rpc_objects"
)

func main() {
	calc := new(rpc_objects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}
/* Output:
Starting Process E:/Go/GoBoek/code_examples/chapter_14/rpc_server.exe ...

** after 5 s: **
End Process exit status 0
*/
