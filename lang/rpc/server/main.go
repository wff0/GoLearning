package main

import (
	rpcdemo "GoBasic/lang/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpcdemo.DemoService{})
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
		}
		//{"method":"abc.def"}
		//{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
		go jsonrpc.ServeConn(conn)
	}
}
