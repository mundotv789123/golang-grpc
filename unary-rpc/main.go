package main

import (
	"flag"
	"fmt"
	"unary-rpc/client"
	"unary-rpc/server"
)

func main() {
	mode := flag.String("mode", "", "modo grpc")
	flag.Parse()

	switch *mode {
	case "server":
		server.Run()
	case "client":
		client.Run()
	default:
		fmt.Println("Escolha o modo correto: server|client")
	}
}
