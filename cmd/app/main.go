package main

import (
	"fmt"

	"github.com/luiztapioca/go-echo-server/pkg/server"
)

func main() {
	fmt.Println("Hello, world!")

	go server.Listen(":8000")

	server.Dial("localhost:8000")
}
