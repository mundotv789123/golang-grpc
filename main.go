package main

import (
	"flag"
)

func main() {
	mode := flag.String("mode", "server", "Modo grpc")
	flag.Parse()
	switch *mode {
	case "server":
		runServer()
	case "client":
		runCliente()
	}
}

func runServer() {

}

func runCliente() {

}
