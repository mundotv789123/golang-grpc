package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
)

type Exemplo struct{}

type Args struct {
	Name string
}

type Reply struct {
	Result string
}

func (*Exemplo) Hello(args *Args, reply *Reply) error {
	reply.Result = fmt.Sprintf("Olá, %s", args.Name)
	return nil
}

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
	exemplo := new(Exemplo)
	rpc.Register(exemplo)
	rpc.HandleHTTP()

	port := 5000
	fmt.Printf("Servidor rodando na porta: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatal(err)
	}
}

func runCliente() {
	client, err := rpc.DialHTTP("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Println("Conectado")

	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		args := Args{Name: text}
		var result Reply
		err = client.Call("Exemplo.Hello", args, &result)
		if err != nil {
			log.Fatal("Error: ", err)
		}

		fmt.Printf("result: %s\n", result.Result)
	}
}
