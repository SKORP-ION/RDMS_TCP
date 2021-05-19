package main

import (
	"RDMS_TCP/handler"
	"RDMS_TCP/utils"
	"log"
	"net"
)

func main() {
	addr := utils.GetServerAddress()

	socket, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("Can't start server. Error:\n", err)
	}

	for {
		conn, err := socket.Accept()

		handler.Accept(conn, err)

	}
}
