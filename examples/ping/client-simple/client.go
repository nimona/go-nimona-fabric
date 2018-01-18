package main

import (
	"context"
	"crypto/tls"
	"fmt"

	fabric "github.com/nimona/go-nimona-fabric"
	ping "github.com/nimona/go-nimona-fabric/examples/ping"
)

func main() {
	crt, err := ping.GenX509KeyPair()
	if err != nil {
		fmt.Println("Cert creation error", err)
		return
	}

	yamux := &fabric.YamuxMiddleware{}
	ident := &fabric.IdentityMiddleware{Local: "CLIENT"}
	security := &fabric.SecMiddleware{
		Config: tls.Config{
			Certificates:       []tls.Certificate{crt},
			InsecureSkipVerify: true,
		},
	}

	f := fabric.New()
	f.AddTransport("tcp", fabric.NewTransportTCP())
	f.AddNegotiatorFunc("yamux", yamux.Negotiate)
	f.AddNegotiatorFunc("tls", security.Negotiate)
	f.AddNegotiatorFunc("identity", ident.Negotiate)

	// make a new connection to the the server's ping handler
	ctx, conn, err := f.DialContext(context.Background(), "tcp:127.0.0.1:3000/tls/yamux/identity/ping")
	if err != nil {
		fmt.Println("Dial error", err)
		return
	}

	// close conection when done
	defer conn.Close()

	rp, ok := ctx.Value(fabric.ContextKeyRemoteIdentity).(string)
	if !ok {
		fmt.Println("Could not find remote id")
		return
	}

	// send ping
	fmt.Println("Ping: Writing ping to", rp)
	if err := fabric.WriteToken(conn, []byte("PING")); err != nil {
		fmt.Println("Could not ping", err)
		return
	}

	fmt.Println("Ping: Wrote ping")

	// get pong
	fmt.Println("Ping: Reading pong...")
	pong, err := fabric.ReadToken(conn)
	if err != nil {
		fmt.Println("Could not read remote pong", err)
		return
	}

	fmt.Println("Ping: Read pong:", string(pong))
}
