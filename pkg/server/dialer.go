package server

import (
	"fmt"
	"log"
	"net"
)

func Dial(port string) {
	conn, err := net.Dial("tcp", port)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("connected to '%s'\n", port)

	_, err = conn.Write([]byte("Hello, world!"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sent")
}
