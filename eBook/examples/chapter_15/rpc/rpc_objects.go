// rpc_objects.go
package rpc_objects

import "net"

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) net.Error {
	*reply = args.N * args.M
	return nil
}