package main

import (
	"bufio"
	"log"
	"net"
	"net/rpc"
	"os"
)

func closeConn(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection closed")
}

func attachConnection(l net.Listener) {
	conn, err := l.Accept()
	defer closeConn(conn)
	if err != nil {
		log.Printf("Errore %+v", err)
	}
	rpc.ServeConn(conn)

}

func connect() {
	reader := bufio.NewReader(os.Stdin)

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Printf("Error while starting rpc server: %+v", err)
	}

	log.Printf("Starting RPC accept...")
	log.Printf("Listening on port 8880 for RPC request")

	for {
		go attachConnection(l)
		text, _ := reader.ReadString('\n')
		if text[0] == 'q' {
			break
		}
	}

}

func Server() {
	api := new(Ret)

	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error in registration!")
	}

	connect()
}
