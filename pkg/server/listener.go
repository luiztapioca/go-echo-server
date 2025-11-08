package server

import (
	"fmt"
	"log"
	"net"
)

func Listen(port string) {
	ln, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Printf("listening on port '%s'\n", port)
		conn, err := ln.Accept()

		if err != nil {
			log.Fatalf("Erro ao aceitar conexao: %v", err)
		}
		go func(c net.Conn) {
			buffer := make([]byte, 1024)
			n, err := c.Read(buffer)
			if err != nil {
				log.Fatal(err)
			}
			if n > 0 {
				fmt.Println(string(buffer[:n]))
			}
		}(conn)
	}
}
