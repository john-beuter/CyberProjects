package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":222")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	response := []byte("Hello from the server")
	_, err = connection.Write(response)

	if err != nil {
		fmt.Println("write error", err) // this is how we can print the var
		return
	}

	connection.Close()

}
