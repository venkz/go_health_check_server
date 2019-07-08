package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	println("Init: Starting Health check server..")
	protocol := flag.String("protocol", "foo", "Protocol for health check server [http | tcp]")
	port := flag.Int("port", -1, "http/tcp port to server health")

	flag.Parse()
	
	if *protocol == "foo" || *port == -1 {
		println("Error: protocol and port are required arguments to start the health server. " +
			"Please refer to ReadMe file for more details.")
		os.Exit(2)
	}

	if *protocol == "http" {
		println(fmt.Sprintf("Server(%s) up and running on port %d", *protocol, *port))
		startHttpServer(*port)
	} else if *protocol == "tcp" {
		println(fmt.Sprintf("Server(%s) up and running on port %d", *protocol, *port))
		startTCPServer(*port)
	} else {
		println("Error Starting server! Only tcp and http protocols are supported")
		os.Exit(2)
	}
}

func startHttpServer(port int) {
	http.HandleFunc("/health", func(writer http.ResponseWriter, reader *http.Request) {
		writer.Header().Set("Server", "Alive and Kicking!")
		writer.WriteHeader(200)
	})

	http.ListenAndServe(fmt.Sprintf("%s:%d", "0.0.0.0", port), nil)
}

func startTCPServer(port int) {
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", "0.0.0.0", port))
	handleError(err)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		handleError(err)

		conn.Write([]byte("Alive and Kicking!"))
		conn.Close()
	}
}

func handleError(err error) {
	if err != nil {
		println("Failed to start a TCP listener!!")
		os.Exit(2)
	}
	return
}
