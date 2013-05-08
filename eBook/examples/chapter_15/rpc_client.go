// rpc_client.go
// if the server is not started:
// can't get the server to start, so client stops immediately with error:
// 2011/08/01 16:08:05 Error dialing:dial tcp :1234: 
//		The requested address is not valid in its context.
// with serverAddress = localhost:
// 2011/08/01 16:09:23 Error dialing:dial tcp 127.0.0.1:1234: 
//		No connection could be made because the target machine actively refused it.
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"./rpc_objects"
)

const serverAddress = "localhost"

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := &rpc_objects.Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
/* Output:
Starting Process E:/Go/GoBoek/code_examples/chapter_14/rpc_client.exe ...

Args: 7 * 8 = 56
End Process exit status 0
*/
