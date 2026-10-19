package main

import (
	"flag"
	"fmt"
	"server-streaming-rpc/client"
	"server-streaming-rpc/server"
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
