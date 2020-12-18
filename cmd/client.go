package main

import (
	"fmt"
	"log"
	"net/rpc"

	"nimona.io/pkg/daemon"
	"nimona.io/pkg/object"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:9000")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	req := daemon.PutObjectRequest{
		// Ctx: context.Background(),
		Obj: &object.Object{
			Type: "lala",
			Data: make(map[string]interface{}),
		},
	}
	res := &daemon.PutObjectResponse{}
	err = client.Call("RPCServer.PutObject", req, res)
	if err != nil {
		log.Fatal("PutObject error:", err)
	}
	fmt.Printf("PutObject: %s --> %s\n", req.Obj.Type, res.Obj.Type)
}
